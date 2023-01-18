package main

import (
	"PartnerPal/configs"
	logger "PartnerPal/logs"
	"PartnerPal/routers"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	configs.ParseConfig("./configs/cfg.yml") //配置读取
	//daos.InitMysql()                         //数据库
	//cache.NewRedis()                         //缓存
	logger.InitLog() //日志
}
func quit() {
	logs.Error("HTTP服务已经优雅关闭")
}
func main() {
	defer func() { //优雅退出
		quit()
	}()
	// 创建一个无中间件的路由
	r := gin.New()
	if gin.Mode() == gin.DebugMode { //判断是否启用日志
		// 启用日志打印
		r.Use(gin.Logger())
	}
	routers.InitRouter(r)                                 //初始化路由
	address := fmt.Sprintf("%s:%d", "0.0.0.0", 8888)      //拼接监听地址
	logs.Info("HTTP服务将在以下地址开启监听：\"https://%s\"", address) //记录启动日志
	if err := r.Run(address); err != nil {                //启动服务
		logs.Error("HTTP服务启动失败")
	}

	sig := make(chan os.Signal) //监听信号量
	signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		for s := range sig {
			switch s {
			case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP:
				quit()
				if i, ok := s.(syscall.Signal); ok {
					os.Exit(int(i))
				} else {
					os.Exit(0)
				}
			}
		}
	}()
	wait := make(chan bool) //等待5秒
	go func() {
		for {
			time.Sleep(5000 * time.Millisecond)
			close(wait)
		}
	}()
	<-wait
}
