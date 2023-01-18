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

// MessageData 消息数据返回结构
type MessageData struct {
	Code uint32 `json:"code"` // 状态码
	Msg  string `json:"msg"`  // 状态描述
}

// JsonMessage 消息返回
func JsonMessage(ctx *gin.Context, code uint32, msg string) {
	ctx.Status(http.StatusOK)
	ctx.JSON(http.StatusOK, Data{
		Code: code,
		Msg:  msg,
	})
	ctx.Abort()
}

// JsonSuccess 成功结果返回
func JsonSuccess(ctx *gin.Context, data interface{}) {
	ctx.Status(http.StatusOK)
	ctx.JSON(200, Data{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
	ctx.Abort()
}

// JsonFailData 失败带结果返回
func JsonFailData(ctx *gin.Context, code uint32, data interface{}) {
	ctx.Status(http.StatusOK)
	ctx.JSON(http.StatusOK, Data{
		Code: code,
		Msg:  "success",
		Data: data,
	})
	ctx.Abort()
}
