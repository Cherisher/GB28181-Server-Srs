package controller

import (
	"net/http"

	"github.com/Cherisher/GB28181-Server-Srs/model"
	"github.com/gin-gonic/gin"
)

//Request URL: http://10.12.30.18:10000/api/v1/stream/start?serial=34220000001320000103&code=44122500141321000002&_=1612702771626

// StreamStart
func StreamStart(ctx *gin.Context) {
	stream := &model.Stream{
		DeviceID:  ctx.Query("serial"),
		ChannelID: ctx.Query("code"),
	}

	if stream.ChannelID == "" || stream.DeviceID == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "参数错误",
		})
		return
	}

	err := stream.Start()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "参数错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, stream)
}

// StreamStop
func StreamStop(ctx *gin.Context) {
	stream := &model.Stream{
		DeviceID:  ctx.Query("serial"),
		ChannelID: ctx.Query("code"),
	}

	if stream.ChannelID == "" || stream.DeviceID == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "参数错误",
		})
		return
	}

	err := stream.Stop()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "参数错误",
		})
		return
	}

	//ctx.JSON(http.StatusOK, stream)
	ctx.String(http.StatusOK, "OK")
}
