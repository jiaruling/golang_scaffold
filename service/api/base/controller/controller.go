package controller

import (
	"GolangScaffold/global"
	"GolangScaffold/service/api/base/dao"
	"GolangScaffold/service/public"
	"github.com/gin-gonic/gin"
	"os"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/19 16:46
*/

func NotFound(c *gin.Context) {
	public.Handler404(c)
	return
}

// ListPage godoc
// @Summary 健康检查
// @Description 健康检查
// @Tags Health
// @ID /demo/bind
// @Accept  json
// @Produce  json
// @Param polygon body dto.DemoInput true "body"
// @Success 200 {object} middleware.Response{data=dto.DemoInput} "success"
// @Router /demo/bind [post]
func Health(c *gin.Context) {
	health := dao.Health{
		Name:      global.ServerConfig.S.Name,
		StartTime: global.StartTime,
		Version:   os.Getenv(global.ServerConfig.S.SystemVersion),
	}
	public.Handler200(c, health)
	return
}

// ListPage godoc
// @Summary ping
// @Description ping
// @Tags Health
// @ID /demo/bind
// @Accept  json
// @Produce  json
// @Param polygon body dto.DemoInput true "body"
// @Success 200 {object} middleware.Response{data=dto.DemoInput} "success"
// @Router /demo/bind [post]
func Ping(c *gin.Context) {
	public.Handler200(c, global.Pong)
	return
}

// redis测试
//func Test(c *gin.Context) {
//	key := "a" + ":" + time.Now().Format("20060102")
//	fmt.Println(key)
//	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
//	defer cancel()
//	if global.REDIS.Exists(ctx, key).Val() == 0 {
//		global.REDIS.HMSet(ctx, key, "issue", 1, "in_hand", 0, "resolve", 0)
//	} else {
//		global.REDIS.HIncrBy(ctx, key, "issue", 1)
//	}
//	public.Handler200(c, nil)
//	return
//}

// 二进制文件转base64后传输
//func GetConfigFile(c *gin.Context) {
//	file := c.DefaultQuery("file", "")
//	path := ""
//	if strings.HasPrefix(file, "/") {
//		path = global.ConfigFilePath + file
//	} else {
//		path = global.ConfigFilePath + "/" + file
//	}
//	if !utils.Exists(path) {
//		public.Handler400(c, errCode.FileNotFound, nil)
//		return
//	}
//	content, err := ioutil.ReadFile(path)
//	if err != nil {
//		public.Handler500(c, errCode.FileReadFailed, nil)
//		return
//	}
//	public.Handler200(c, base64.StdEncoding.EncodeToString(content))
//	return
//}

// base64数据解码后写入文件
//func PostConfigFile(c *gin.Context) {
//	var v dto.FileContent
//	if err := c.ShouldBind(&v); err != nil {
//		public.FormsVerifyFailed(c, err)
//		return
//	}
//	content, err := base64.StdEncoding.DecodeString(v.Content)
//	if err != nil {
//		public.Handler500(c, errCode.Base64DecodeFailed, nil)
//		return
//	}
//	if err := ioutil.WriteFile(global.ConfigFilePath+v.FileName, content, os.ModePerm); err != nil {
//		public.Handler500(c, errCode.FileWriteFailed, nil)
//		return
//	}
//	public.Handler200(c, nil)
//	return
//}

// https://github.com/oschwald/geoip2-golang
//func IpToCountry(c *gin.Context) {
//
//}
