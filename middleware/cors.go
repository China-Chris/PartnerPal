package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// Cors 跨域支持
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,                                                                                                     // 是否允许全部人访问
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},                                                // 允许的请求类型
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Sign", "Timestamp", "x-requested-with", "Version"}, // 允许的请求头部
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},                                                               // 暴露出去的请求头
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
}
