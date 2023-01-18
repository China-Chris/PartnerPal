package activeCount

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// DailyActiveCount 日活统计中间件
func DailyActiveCount(ctx gin.Context) {
	//获取当前日期
	now := time.Now()
	data := fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
	//在Redis中计数器+1
	fmt.Sprintf(data)
	ctx.Next()
}
