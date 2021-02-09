package controller

import (
	"net/http"
	"time"

	"github.com/Cherisher/GB28181-Server-Srs/model"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	user := &model.User{
		Username: ctx.Query("username"),
		Password: ctx.Query("password"),
	}

	if user.Username == "" || user.Password == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "输入参数错误",
		})
		return
	}

	if ctx.DefaultQuery("url_token_only", "false") == "true" {
		user.URLTokenOnly = true
	}

	ok := user.Verify()
	if ok {
		if user.URLTokenOnly {
			//将不会入库持久 session, 减少库表操作
		}

		type authResp struct {
			URLToken     string `json:"URLToken,omitempty"`    //用于方式一的鉴权
			CookieToken  string `json:"CookieToken,omitempty"` //用于方式二的鉴权
			TokenTimeout int    `json:"TokenTimeout"`          //Token 超时(秒)

			AuthToken string `json:"AuthToken,omitempty"` //等同于 URLToken, 为了兼容保留
			Token     string `json:"Token,omitempty"`     //等同于 CookieToken, 为了兼容保留
		}

		ctx.JSON(http.StatusOK, authResp{
			URLToken:     "test.token.12345",
			TokenTimeout: 3600,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1002,
			"msg":  "鉴权失败",
		})
	}

	// rsp := &loginRsp{
	// 	URLToken:     "slkdfjiw2ej.jskdfwe-wefadsf.sdfwoeuyqtew",
	// 	CookieToken:  "mDC4tu-ig",
	// 	TokenTimeout: 604800,
	// }
	// ctx.JSON(200, rsp)
}

func ServerInfo(ctx *gin.Context) {
	type serverInfoRsp struct {
		APIAuth          bool   `json:"APIAuth"`
		Authorization    string `json:"Authorization"`
		ChannelCount     int    `json:"ChannelCount"`
		CopyrightText    string `json:"CopyrightText"`
		Hardware         string `json:"Hardware"`
		InterfaceVersion string `json:"InterfaceVersion"`
		IsDemo           bool   `json:"IsDemo"`
		LogoMiniText     string `json:"LogoMiniText"`
		LogoText         string `json:"LogoText"`
		RemainDays       int    `json:"RemainDays"`
		RunningTime      string `json:"RunningTime"`
		SIPHost          string `json:"SIPHost"`
		SIPPort          int    `json:"SIPPort"`
		SIPRealm         string `json:"SIPRealm"`
		SIPSerial        string `json:"SIPSerial"`
		Server           string `json:"Server"`
		ServerTime       string `json:"ServerTime"`
		ShowAbout        bool   `json:"ShowAbout"`
		ShowShare        bool   `json:"ShowShare"`
		StartUpTime      string `json:"StartUpTime"`
		VersionType      string `json:"VersionType"`
	}

	rsp := &serverInfoRsp{
		APIAuth:          false,
		Authorization:    "Users",
		ChannelCount:     100,
		CopyrightText:    "Copyright © 2020 \u003ca href=\"//www.videiot.cn\" target=\"_blank\"\u003ewww.videiot.cn\u003c/a\u003e All rights reserved.",
		Hardware:         "AMD64",
		InterfaceVersion: "v1",
		IsDemo:           false,
		LogoMiniText:     "GBS",
		LogoText:         "CmiotGBS",
		RemainDays:       1000,
		RunningTime:      "19 Days 20 Hours 11 Mins 52 Secs",
		SIPHost:          "10.12.30.20",
		SIPPort:          15160,
		SIPRealm:         "3412000000",
		SIPSerial:        "34120000002000000001",
		Server:           "IotGBS/v4.11",
		ServerTime:       time.Now().Format("2006-01-02 15:04:05"),
		ShowAbout:        true,
		ShowShare:        true,
		StartUpTime:      "2021-01-18 17:53:44",
		VersionType:      "旗舰版",
	}
	ctx.JSON(200, rsp)
}

func Userinfo(ctx *gin.Context) {
	type userInfoRsp struct {
		ID            int      `json:"ID"`
		Name          string   `json:"Name"`
		Roles         []string `json:"Roles"`
		HasAllChannel bool     `json:"HasAllChannel"`
	}

	rsp := &userInfoRsp{
		ID:            1,
		Name:          "admin",
		Roles:         []string{"超级管理员"},
		HasAllChannel: true,
	}

	ctx.JSON(200, rsp)
}

func DashboardTop(ctx *gin.Context) {
	type Data struct {
		CPU    []*model.ResourceUsage `json:"cpuData"`
		Memory []*model.ResourceUsage `json:"memData"`
	}

	type topRsp struct {
		Code int    `json:"code"`
		Data Data   `json:"data"`
		Msg  string `json:"msg"`
	}

	rsp := &topRsp{
		Code: 200,
		Msg:  "Success",
		Data: Data{
			CPU:    model.CPUUsage(),
			Memory: model.MemUsage(),
		},
	}

	ctx.JSON(200, rsp)
}

// DashboardStore 存储使用
func DashboardStore(ctx *gin.Context) {

	type storeRsp struct {
		Code int               `json:"code"`
		Data []*model.DiskInfo `json:"data"`
		Msg  string            `json:"msg"`
	}

	rsp := &storeRsp{
		Code: 200,
		Data: model.DiskStore(),
	}

	ctx.JSON(200, rsp)
}

// DashboardAuth 获取通道信息
func DashboardAuth(ctx *gin.Context) {
	type Data struct {
		ChannelCount  int `json:"ChannelCount"`
		ChannelOnline int `json:"ChannelOnline"`
		ChannelTotal  int `json:"ChannelTotal"`
		DeviceOnline  int `json:"DeviceOnline"`
		DeviceTotal   int `json:"DeviceTotal"`
	}

	type authRsp struct {
		Code int    `json:"code"`
		Data Data   `json:"data"`
		Msg  string `json:"msg"`
	}

	devicesOnlineCount := 0
	devicesTotalCount := 0
	channelOnlineCount := 0
	channelTotalCount := 0

	devices, _ := model.DeviceList()
	devicesTotalCount = len(devices)
	for _, device := range devices {
		if device.Online == true {
			devicesOnlineCount++
		}

		channels, _ := device.ChannelList()
		channelTotalCount += len(channels)

		for _, channel := range channels {
			if channel.Status == "ON" {
				channelOnlineCount++
			}
		}
	}

	rsp := &authRsp{
		Code: 200,
		Msg:  "Success",
		Data: Data{
			ChannelCount:  0, // 含义未知
			ChannelOnline: channelOnlineCount,
			ChannelTotal:  channelTotalCount,
			DeviceOnline:  devicesOnlineCount,
			DeviceTotal:   devicesTotalCount,
		},
	}

	ctx.JSON(200, rsp)
}
