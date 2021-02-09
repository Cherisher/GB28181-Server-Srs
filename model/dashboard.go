package model

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

type resultCPU struct {
	id         int64
	moduleName string
	success    bool
	time       time.Time
	percent    float64
}

// resultHistory contains two history slices: `results` contains most recent `maxResults` results.
// After they expire out of `results`, failures will be saved in `preservedFailedResults`. This
// ensures that we are always able to see debug information about recent failures.
type resultHistory struct {
	mu                     sync.Mutex
	nextID                 int64
	results                []*resultCPU
	preservedFailedResults []*resultCPU
	maxResults             uint
}

// Add a result to the history.
func (rh *resultHistory) Add(moduleName string, percent float64, success bool) {
	rh.mu.Lock()
	defer rh.mu.Unlock()

	r := &resultCPU{
		id:         rh.nextID,
		moduleName: moduleName,
		time:       time.Now(),
		percent:    percent,
		success:    success,
	}
	rh.nextID++

	rh.results = append(rh.results, r)
	if uint(len(rh.results)) > rh.maxResults {
		// If we are about to remove a failure, add it to the failed result history, then
		// remove the oldest failed result, if needed.
		if !rh.results[0].success {
			rh.preservedFailedResults = append(rh.preservedFailedResults, rh.results[0])
			if uint(len(rh.preservedFailedResults)) > rh.maxResults {
				preservedFailedResults := make([]*resultCPU, len(rh.preservedFailedResults)-1)
				copy(preservedFailedResults, rh.preservedFailedResults[1:])
				rh.preservedFailedResults = preservedFailedResults
			}
		}
		results := make([]*resultCPU, len(rh.results)-1)
		copy(results, rh.results[1:])
		rh.results = results
	}
}

// List returns a list of all results.
func (rh *resultHistory) List() []*resultCPU {
	rh.mu.Lock()
	defer rh.mu.Unlock()

	// Results in each slice are disjoint. We can simply concatenate the results.
	return append(rh.preservedFailedResults[:], rh.results...)
}

// Get returns a given result.
func (rh *resultHistory) Get(id int64) *resultCPU {
	rh.mu.Lock()
	defer rh.mu.Unlock()

	for _, r := range rh.preservedFailedResults {
		if r.id == id {
			return r
		}
	}
	for _, r := range rh.results {
		if r.id == id {
			return r
		}
	}

	return nil
}

var (
	cpuLoad *resultHistory
	memInfo *resultHistory
)

// ResourceUsage CPU OR Memory
type ResourceUsage struct {
	Time string  `json:"time"`
	Use  float64 `json:"use"`
}

// CPUUsage xx
func CPUUsage() []*ResourceUsage {
	savedResults := cpuLoad.List()
	if len(savedResults) == 0 {
		return nil
	}
	datas := make([]*ResourceUsage, 0, len(savedResults))
	for _, r := range savedResults {
		datas = append(datas, &ResourceUsage{
			Time: r.time.Format(timeFormart),
			Use:  r.percent,
		})
	}
	return datas
}

// MemUsage 内存占用
func MemUsage() []*ResourceUsage {
	savedResults := memInfo.List()
	if len(savedResults) == 0 {
		return nil
	}
	datas := make([]*ResourceUsage, 0, len(savedResults))
	for _, r := range savedResults {
		datas = append(datas, &ResourceUsage{
			Time: r.time.Format(timeFormart),
			Use:  r.percent,
		})
	}
	return datas
}

func taskFunc() {
	// cpuInfos, err := cpu.Info()
	// if err != nil {
	// 	fmt.Printf("get cpu info failed, err:%v", err)
	// }
	// for _, ci := range cpuInfos {
	// 	fmt.Println(ci)
	// }
	// CPU使用率
	// for {
	// 	percent, _ := cpu.Percent(time.Second, false)
	// 	fmt.Printf("cpu percent:%v\n", percent)
	// }
	//percent, _ := cpu.Percent(time.Second, false)
	//fmt.Printf("%v cpu percent:%v\n", time.Now(), percent)
	//获取CPU负载信息
	cInfo, _ := load.Avg()
	cpuLoad.Add("cpu", cInfo.Load1/100.0, true)
	// Memory
	mInfo, _ := mem.VirtualMemory()
	memInfo.Add("memory", mInfo.UsedPercent/100.0, true)
}

// DiskInfo 存储使用
type DiskInfo struct {
	Name      string `json:"Name"`
	Unit      string `json:"Unit"`
	Size      string `json:"Size"`
	FreeSpace string `json:"FreeSpace"`
	Used      string `json:"Used"`
	Percent   string `json:"Percent"`
	Threshold string `json:"Threshold"`
}

// DiskStore disk
func DiskStore() []*DiskInfo {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return nil
	}

	if len(parts) == 0 {
		return nil
	}

	datas := make([]*DiskInfo, 0, len(parts))

	for _, part := range parts {
		//fmt.Println(part.String())
		matched, err := regexp.Match(`^/dev/[sx]d[a-z][0-9]*`, []byte(part.Device))
		if err != nil || matched == false {
			continue
		}
		if u, err := disk.Usage(part.Mountpoint); err == nil && u.Total > 0 {
			datas = append(datas, &DiskInfo{
				Name:      part.Mountpoint,
				Unit:      "G",
				Size:      strconv.FormatFloat(float64(u.Total)/1024.0/1024.0/1024.0, 'f', 2, 64),
				FreeSpace: strconv.FormatFloat(float64(u.Free)/1024.0/1024.0/1024.0, 'f', 2, 64),
				Used:      strconv.FormatFloat(float64(u.Used)/1024.0/1024.0/1024.0, 'f', 2, 64),
				Percent:   strconv.FormatFloat(float64(u.UsedPercent), 'f', 2, 64),
			})
		}
	}
	return datas
}

func init() {
	cpuLoad = &resultHistory{maxResults: 10}
	memInfo = &resultHistory{maxResults: 10}

	go func(interval int, function func()) {
		eventsTick := time.NewTicker(time.Duration(interval) * time.Second)
		defer func() {
			eventsTick.Stop()
		}()

		for {
			select {
			case <-eventsTick.C:
				function()
			}
		}
	}(1 /*单位:second*/, taskFunc)
}
