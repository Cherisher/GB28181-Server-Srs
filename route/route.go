package route

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Cherisher/GB28181-Server-Srs/controller"
	"github.com/gin-gonic/gin"
)

// Router 全局路由
var Router *gin.Engine

func init() {
	Router = gin.New()

	Router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	Router.Use(gin.Recovery())

}

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	Router.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "The incorrect API route.")
	})

	v1 := Router.Group("/api/v1")
	{
		// 系统接口
		system := v1
		{
			system.GET("/login", controller.Login)
			system.GET("/getserverinfo", controller.ServerInfo)
			system.GET("/userinfo", controller.Userinfo)
		}

		dashboard := v1.Group("/dashboard")
		{
			dashboard.GET("/top", controller.DashboardTop)
			dashboard.GET("/store", controller.DashboardStore)
			dashboard.GET("/auth", controller.DashboardAuth)
		}

		device := v1.Group("/device")
		{
			device.GET("/list", controller.DeviceList)
			device.GET("/info", controller.DeviceInfo)
			device.GET("/fetchcatalog", controller.DeviceFetchCatalog)

			device.GET("/channellist", controller.DeviceChannelList)
		}

		stream := v1.Group("/stream")
		{
			stream.GET("/start", controller.StreamStart)
		}

	}
	return Router
}
