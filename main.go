package main

import (
	"GolangScaffold/global"
	"GolangScaffold/initial"
	"github.com/jiaruling/golang_utils/lib"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	var (
		err  error
		quit chan os.Signal
	)
	// 初始化项目目录
	lib.InitDir([]string{global.LogFilePath, global.FILEPATH}, []string{"log"})
	// 加载配置文件
	initial.InitConfigFile()
	// 初始化日志配置
	lib.InitLog(global.ServerConfig.L.LogFileDir, global.ServerConfig.L.AppName, global.ServerConfig.L.MaxSize,
		global.ServerConfig.L.MaxBackups, global.ServerConfig.L.MaxAge, global.ServerConfig.L.Dev)
	// 获取日志
	log := lib.GetLog()
	trace := lib.NewTrace()
	// 线程数等于cpu核数
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 初始化数据库
	initial.InitDB()
	// 初始化翻译器
	initial.InitTrans("zh") // zh or en
	// 初始化验证器
	if err = initial.InitValidator(); err != nil {
		log.Panic(trace, lib.DLTagUndefind, map[string]interface{}{"hint": "加载配置文件失败", "error": err.Error()})
		os.Exit(1)
	}
	initial.InitGin()
	initial.InitRouter()
	initial.InitService(trace)
	initial.InitTask()

	log.Info(trace, lib.DLTagUndefind, map[string]interface{}{
		"hint": "服务启动成功",
		"cpuNum": runtime.NumCPU(),
		"GOOS": runtime.GOOS,
	})
	// 优雅退出
	quit = make(chan os.Signal) // 定义一个无缓冲的通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	lib.DestroyMysqlGormAll()
	lib.DestroyRedisAll()
	log.Info(trace, lib.DLTagUndefind, map[string]interface{}{
		"hint": "优雅退出",
		"exit1": "关闭mysql连接",
		"exit2": "关闭redis连接",
	})
}
