package initial

import (
	"GolangScaffold/global"
	"fmt"
	"github.com/jiaruling/golang_utils/lib"
)

func InitDB() {
	// Mysql
	for _, m := range global.ServerConfig.M {
		mysql := lib.NewMysqlGorm(m.User, m.Password, m.Ip, m.Port, m.DB)
		mysql.Name = m.Name
		mysql.Parameter = m.Parameter
		mysql.InitMysqlGorm()
	}
	// Redis
	for _, r := range global.ServerConfig.R {
		redis := lib.NewRedis(fmt.Sprintf("%s:%v", r.Ip, r.Port), r.Password, r.DB)
		redis.Name = r.Name
		redis.InitRedis()
	}
}