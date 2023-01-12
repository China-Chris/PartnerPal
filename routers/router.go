package routers

import (
	"PartnerPal/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

// InitRouter 初始化路由
func InitRouter(router *gin.Engine) {
	// 为所有路由添加Recover异常捕获（返回自己定义的内容）
	router.Use(middleware.AppRecover)
	// 注册自定义全局跨域支持
	router.Use(middleware.Cors())
	// 定义根路径服务启动状态检测
	router.HEAD("/head", func(ctx *gin.Context) {
		//tools.JsonSuccess(ctx, nil)
	})
	docs.SwaggerInfo.BasePath = "/api"
	// 定义api路由分组
	//nft := router.Group("/nft")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
