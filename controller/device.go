package controller

import (
	"net/http"
	"strconv"

	"github.com/Cherisher/GB28181-Server-Srs/model"
	"github.com/gin-gonic/gin"
)

//DeviceList 查询设备列表
func DeviceList(ctx *gin.Context) {
	type query struct {
	}
	type deviceListRsp struct {
		DeviceCount int             `json:"DeviceCount"`
		DeviceList  []*model.Device `json:"DeviceList"`
	}
	var rsp deviceListRsp

	devices, err := model.DeviceList()

	if err != nil {
		rsp.DeviceCount = 0
		rsp.DeviceList = make([]*model.Device, 0)
		ctx.JSON(200, rsp)
		return
	}

	rsp.DeviceCount = len(devices)
	rsp.DeviceList = devices

	ctx.JSON(200, rsp)

}

//DeviceInfo 查询单条设备信息
func DeviceInfo(ctx *gin.Context) {
	type query struct {
		Serial string
	}

	q := query{
		Serial: ctx.Query("serial"),
	}

	if q.Serial == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "参数错误",
		})
		return
	}

	device := &model.Device{
		ID: q.Serial,
	}

	if err := device.DeviceInfo(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1002,
			"msg":  "数据查询错误[" + q.Serial + "]:" + err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, device)
	}

}

// DeviceFetchCatalog 获取下级通道列表
func DeviceFetchCatalog(ctx *gin.Context) {
	type query struct {
		Serial  string
		Timeout int
	}
	q := query{
		Serial: ctx.Query("serial"),
	}

	q.Timeout, _ = strconv.Atoi(ctx.DefaultQuery("timeout", "15"))

	type catalogResp struct {
		ChannelCount int                   `json:"ChannelCount"`
		ChannelList  []model.DeviceChannel `json:"ChannelList"`
	}

	resp := catalogResp{
		ChannelCount: 1,
		ChannelList: []model.DeviceChannel{
			0: {
				DeviceID: "34220000001320000103",
				Name:     "测试通道",
				PTZType:  3,
				Parental: 0,
			},
		},
	}

	ctx.JSON(http.StatusOK, &resp)
}

//DeviceChannelList
func DeviceChannelList(ctx *gin.Context) {
	type query struct {
		Serial      string
		ChannelType string // 通道类型, device - 直接子设备, dir - 直接子目录, all - 所有默认值: all
		DirSerial   string //子目录编号
		Start       int    //分页开始,从零开始
		Limit       int    //分页大小
		Q           string //搜索关键字
		Online      bool   //在线状态		允许值: true, false
	}

	q := query{
		Serial:      ctx.Query("serial"),
		ChannelType: ctx.DefaultQuery("channel_type", "all"),
	}

	type channelResp struct {
		ChannelCount int                    `json:"ChannelCount"`
		ChannelList  []*model.DeviceChannel `json:"ChannelList"`
	}

	//fmt.Println(q.Serial)

	dev := &model.Device{
		ID: q.Serial,
	}

	channels, err := dev.ChannelList()
	if err != nil {
		//fmt.Println("dev.ChannelList error:", err.Error())
		ctx.JSON(http.StatusOK, &channelResp{
			ChannelCount: 0,
			ChannelList:  make([]*model.DeviceChannel, 0),
		})
		return
	}
	rsp := &channelResp{
		ChannelCount: len(channels),
		ChannelList:  channels,
	}

	ctx.JSON(http.StatusOK, rsp)

}
