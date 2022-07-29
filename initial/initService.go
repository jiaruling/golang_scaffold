package initial

import (
	"GolangScaffold/global"
	"github.com/jiaruling/golang_utils/lib"
)

/*
   功能说明: 启动服务
   参考:
   创建人: 贾汝凌
   创建时间: 2022/game_base_service.sql/19 14:45
*/

func InitService(trace *lib.TraceContext) {
	go func() {
		log := lib.GetLog()
		if err := global.GinRouter.Router.Run(global.ServerConfig.S.Ip + ":" + global.ServerConfig.S.Port); err != nil {
			log.Panic(trace, lib.DLTagUndefind, map[string]interface{}{
				"hint":  "http服务启动失败",
				"error": err.Error(),
			})
		}
	}()
}
