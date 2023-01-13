package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Data 数据返回结构
type Data struct {
	Code uint32      `json:"code"`           // 状态码
	Msg  string      `json:"msg"`            // 状态描述
	Data interface{} `json:"data,omitempty"` // 常规json数据或PaginationData数据
}

// JsonSuccess 成功结果返回
// @params ctx  *gin.Context Gin框架上下文
// @params data interface{}  返回数据
func JsonSuccess(ctx *gin.Context, data interface{}) {
	ctx.Status(http.StatusOK)
	ctx.JSON(200, Data{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
	ctx.Abort()
}
