package errorss

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
)

const (
	ErrInternalServer = iota + 500 //Post 解析失败 参数错误
)
const (
	ErrShouldBind = iota + 1000 //Post 解析失败 参数错误
	ErrSignUp
	ErrCheckMobile
)

// ErrorCodeTextMap 定义错误码和错误信息的映射表
var ErrorCodeTextMap = map[int]string{
	ErrInternalServer: "内部服务发生错误",
	ErrShouldBind:     "参数解析错误请确认参数正确",
	ErrSignUp:         "注册时发生错误",
	ErrCheckMobile:    "手机号校验不匹配",
}

//HandleCustomError 统一自定义错误处理
func HandleCustomError(code int, msg string, err error) error {
	err = fmt.Errorf("code:%d, %s，err:%s", code, msg, err)
	logs.Error(err)
	return err
}

//HandleError 统一错误处理程序
func HandleError(code int, err error) error {
	return HandleCustomError(code, ErrorCodeTextMap[code], err)
}
