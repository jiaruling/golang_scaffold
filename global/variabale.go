package global

import (
	"GolangScaffold/global/config"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"time"
)

type Router struct {
	Router *gin.Engine
	V1     *gin.RouterGroup
}

var (
	ServerConfig config.ServerConfig
	Trans        ut.Translator
	GinRouter    *Router
	StartTime   string
)

func init() {
	StartTime = time.Now().Format("2006.01.02 15:04:05")
}