package main

import (
	"PartnerPal/configs"
	"PartnerPal/logers"
	"PartnerPal/routers"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.ParseConfig("./configs/cfg.yml") //配置读取
	//daos.InitMysql()                         //数据库
	//cache.NewRedis()                         //缓存
	loger.InitLog() //日志
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth
func main() {
	// 创建一个无中间件的路由
	r := gin.New()
	// 判断是否启用日志
	if gin.Mode() == gin.DebugMode {
		// 启用日志打印
		r.Use(gin.Logger())
	}
	// 初始化路由
	routers.InitRouter(r)
	// 拼接监听地址q
	address := fmt.Sprintf("%s:%d", "0.0.0.0", 9999)
	// 记录启动日志
	logs.Info("HTTP服务将在以下地址开启监听：\"http://%s\"", address)
	// 启动服务
	if err := r.Run(address); err != nil {
		logs.Error("HTTP服务启动失败")
	}
}
