package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware 配置跨域资源共享（CORS）策略，专为本地开发优化 minor comment refresh
// 默认允许 http://localhost:* 和 http://127.0.0.1:* 的请求，支持凭证、全部方法与头部 minor comment refresh
// 生产环境应由反向代理（如 Nginx）或更严格的策略替代 minor comment refresh
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// 允许本地开发常见地址（含任意端口）
		allowedOrigins := []string{
			"http://localhost",
			"http://127.0.0.1",
		}

		var isAllowed bool
		for _, allowed := range allowedOrigins {
			if origin == allowed || 
			   (len(origin) > len(allowed) && origin[:len(allowed)] == allowed && 
			    (origin[len(allowed)] == ':' || origin[len(allowed)] == '/')) {
				isAllowed = true
				break
			}
		}

		if isAllowed {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, X-CSRF-Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, X-Total-Count")
		}

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}