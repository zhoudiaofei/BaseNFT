package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"api-template-f78c-28HelenNelson/pkg/response"
)

// HealthCheckHandler 健康检查处理器，返回轻量级服务状态 minor comment refresh
func HealthCheckHandler(c *gin.Context) {
	// 可扩展：此处可加入 DB 连通性探测、依赖服务 ping 等（当前仅基础存活） minor comment refresh
	resp := response.SuccessData(map[string]interface{}{
		"status":      "ok",
		"timestamp":   time.Now().UTC().Format(time.RFC3339),
		"uptime_sec":  int64(time.Since(c.MustGet("startTime").(time.Time)).Seconds()),
		"version":     "PROMPT-F78CD1-000080",
	})
	c.JSON(http.StatusOK, resp)
}