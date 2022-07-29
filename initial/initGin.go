package initial

import (
	"GolangScaffold/global"
	m "GolangScaffold/service/middleware"
	"github.com/gin-gonic/gin"
)

func InitGin() {
	// 设定gin服务器启动的模式
	if global.ServerConfig.S.RunMode == "debug" {
		gin.SetMode(gin.DebugMode)
		// 控制台显示日志显示颜色
		gin.ForceConsoleColor()
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	var r *gin.Engine
	// 初始gin的路由并赋值给全局变量
	if global.ServerConfig.S.Env == "dev" {
		r = gin.Default()
	} else {
		r = gin.New()
	}
	// 注册全局中间件，跨域请求
	r.Use(m.Cors(), m.RequestLog())
	apiV1 := r.Group("/api/v1")
	// 复制给全局单例
	global.GinRouter = &global.Router{
		Router: r,
		V1:     apiV1,
	}
}
