package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 记录 HTTP 请求的基本信息：方法、路径、状态码、耗时（毫秒） minor comment refresh
// 仅在非测试环境启用，避免干扰单元测试输出 minor comment refresh
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求 minor comment refresh
		c.Next()

		// 计算耗时 minor comment refresh
		latency := time.Since(start)

		// 获取状态码 minor comment refresh
		status := c.Writer.Status()

		// 日志格式：[METHOD] PATH - STATUS - LATENCY
		// 示例：[GET] /health - 200 - 1.234ms
		fmt.Printf("[%s] %s - %d - %.3fms\n",
			c.Request.Method,
			c.Request.URL.Path,
			status,
			latency.Seconds()*1000.0,
		)
	}
}