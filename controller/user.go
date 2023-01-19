package controller

import (
	"PartnerPal/pkg/errorss/errors_const"
	"PartnerPal/pkg/request"
	"PartnerPal/pkg/response"
	"PartnerPal/service/user"
	"github.com/gin-gonic/gin"
)

// Login 用户登陆
func Login(ctx *gin.Context) {
	response.JsonSuccess(ctx, nil)
}

// SignUp 用户注册
func SignUp(ctx *gin.Context) {
	var signUp request.RqSignUp
	err := ctx.ShouldBindJSON(&signUp)
	if err != nil {
		response.JsonFailMessage(ctx, errors_const.ErrInternalServer, err)
		return
	}
	checkMobile := user.CheckMobile(signUp.Phone)
	if !checkMobile {
		response.JsonFailMessage(ctx, errors_const.ErrCheckMobile, err)
		return
	}
	date, err := user.SignUp(signUp)
	if err != nil {
		response.JsonFailMessage(ctx, errors_const.ErrSignUp, err)
		return
	}
	response.JsonSuccess(ctx, date)
}
