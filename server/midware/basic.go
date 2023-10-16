package midware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RunTimeMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Header("test", "test")
		// 继续处理请求
		c.Next()

		// 记录结束时间并计算运行时间
		end := time.Now()
		duration := end.Sub(start)

		// 添加自定义响应头，将运行时间作为响应的一部分
		c.Header("X-RunTime", duration.String())
		c.Header("X-Request-Time", fmt.Sprintf("%.2fms", duration.Seconds()*1000))

	}
}
