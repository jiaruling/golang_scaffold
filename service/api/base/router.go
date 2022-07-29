package base

import (
	"GolangScaffold/global"
	"GolangScaffold/service/api/base/controller"
	"net/http"
)

func Router() {
	//加载静态资源，一般是上传的资源，例如用户上传的图片
	global.GinRouter.Router.StaticFS("/static", http.Dir(global.FILEPATH))
	global.GinRouter.Router.StaticFS("/logs", http.Dir(global.LogFilePath))
	// 404 路由
	global.GinRouter.Router.NoRoute(controller.NotFound)
	// 服务健康检查 eureka服务健康检查
	global.GinRouter.Router.GET("/health", controller.Health)
	global.GinRouter.Router.GET("/info", controller.Health)
	global.GinRouter.Router.GET("/ping", controller.Ping)
	//global.GinRouter.Router.GET("/test", controller.Test)
	// 文件base64获取与上传
	//global.GinRouter.Router.GET("/file/base64", controller.GetConfigFile)
	//global.GinRouter.Router.POST("/file/base64", controller.PostConfigFile)
	// IP地址转换为地区国家
	//global.GinRouter.Router.GET("/ip_to_country", controller.IpToCountry)
}
