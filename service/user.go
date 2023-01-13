package service

import (
	"PartnerPal/pkg/response"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	response.JsonSuccess(ctx, 200)
}
