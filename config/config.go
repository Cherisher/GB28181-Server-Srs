package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

// SrsbaseConfig SRS服务器配置
type SrsbaseConfig struct {
	APIBase  string `json:"api_base"`
	RtmpBase string `json:"rtmp_base"`
	FlvBase  string `json:"flv_base"`
}

// SelfConfig 私有配置
type SelfConfig struct {
	Tag  int `json:"tag"`
	Port int `json:"port"`
}

// GlobalConfig 全局配置
type GlobalConfig struct {
	Srs  SrsbaseConfig `json:"srs"`
	Self SelfConfig    `json:"self"`
}

var (
	globalconfig *GlobalConfig
	configMux    sync.RWMutex
)

// Config 全局配置
func Config() *GlobalConfig {
	return globalconfig
}

// InitConfigJson 初始化配置
func InitConfigJson(fliepath string) error {
	var config GlobalConfig
	file, err := ioutil.ReadFile(fliepath)
	if err != nil {
		fmt.Println("配置文件读取错误,找不到配置文件", err)
		return err
	}

	if err = json.Unmarshal(file, &config); err != nil {
		fmt.Println("配置文件读取失败", err)
		return err
	}

	if config.Self.Port == 0 {
		config.Self.Port = 8080
	}

	configMux.Lock()
	globalconfig = &config
	configMux.Unlock()
	return nil
}
