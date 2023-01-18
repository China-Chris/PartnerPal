package controller

import (
	errorss "PartnerPal/pkg/errors"
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
	var signUp request.SignUp
	err := ctx.ShouldBind(&signUp)
	if err != nil {
		response.JsonFailData(ctx, errorss.ErrInternalServer, errorss.HandleError(errorss.ErrInternalServer, err))
		return
	}
	checkMobile := user.CheckMobile(signUp.Phone)
	if !checkMobile {
		response.JsonFailData(ctx, errorss.ErrInternalServer, errorss.HandleError(errorss.ErrCheckMobile, err))
		return
	}
	_, err = user.SignUp(signUp)
	if err != nil {
		response.JsonFailData(ctx, errorss.ErrInternalServer, errorss.HandleError(errorss.ErrSignUp, err))
		return
	}
	response.JsonSuccess(ctx, nil)
}
