package controller

import (
	"PartnerPal/pkg/response"
	"github.com/gin-gonic/gin"
)

// Login 用户登陆
func Login(ctx *gin.Context) {
	response.JsonSuccess(ctx, 200)
}

// SignUp 用户注册
func SignUp(ctx *gin.Context) {
	response.JsonSuccess(ctx, 200)
}
