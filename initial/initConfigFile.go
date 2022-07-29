package initial

import (
	"GolangScaffold/global"
	"fmt"
	"github.com/jiaruling/golang_utils/lib"
	"os"
)

func InitConfigFile() {
	if err := lib.ParseConfig("./config/server.yml", &global.ServerConfig, false); err != nil {
		fmt.Println("加载配置文件失败001: ", err.Error())
		os.Exit(1)
	}
	return
}
