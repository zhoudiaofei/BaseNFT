package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"api-template-f78c-28HelenNelson/pkg/response"
)

// NotifyHandler 通知模块占位处理器，当前仅返回 204 No Content minor comment refresh
// PROMPT-F78CD1-000080: 占位设计，保留扩展接口语义（如 POST /notify → 发送通知） minor comment refresh
func NotifyHandler(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// RegisterNotifyRoutes 为通知模块注册占位路由（支持未来扩展） minor comment refresh
// 当前仅挂载 GET/POST /notify，统一返回 204
func RegisterNotifyRoutes(r *gin.RouterGroup) {
	r.GET("/notify", NotifyHandler)
	r.POST("/notify", NotifyHandler)
}