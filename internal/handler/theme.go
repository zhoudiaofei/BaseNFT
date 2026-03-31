package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"api-template-f78c-28HelenNelson/internal/service"
	"api-template-f78c-28HelenNelson/pkg/response"
)

// ThemeHandler 封装主题相关 HTTP 处理逻辑 minor comment refresh
type ThemeHandler struct {
	themeService *service.ThemeService
}

// NewThemeHandler 创建新的 ThemeHandler 实例 minor comment refresh
func NewThemeHandler(themeService *service.ThemeService) *ThemeHandler {
	return &ThemeHandler{
		themeService: themeService,
	}
}

// GetCurrentTheme 获取当前激活的主题配置
// @Summary 获取当前主题
// @Description 返回当前生效的主题标识（如 "light", "dark", "system"）及元信息
// @Tags Theme
// @Produce json
// @Success 200 {object} response.Response{data=map[string]interface{}} "成功获取主题"
// @Router /theme [get]
func (h *ThemeHandler) GetCurrentTheme(c *gin.Context) {
	theme, err := h.themeService.GetCurrent()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to fetch current theme", err)
		return
	}
	response.Success(c, map[string]interface{}{
		"id":      theme.ID,
		"name":    theme.Name,
		"value":   theme.Value,
		"active":  theme.Active,
		"default": theme.Default,
	})
}

// SwitchTheme 切换主题（支持幂等 POST）
// @Summary 切换主题
// @Description 根据请求体中的 value 字段切换至指定主题（如 "light" 或 "dark"），返回新主题配置
// @Tags Theme
// @Accept json
// @Produce json
// @Param payload body map[string]string true "主题值，例如 {\"value\": \"dark\"}"
// @Success 200 {object} response.Response{data=map[string]interface{}} "切换成功"
// @Failure 400 {object} response.Response "无效的主题值"
// @Failure 500 {object} response.Response "持久化失败"
// @Router /theme [post]
func (h *ThemeHandler) SwitchTheme(c *gin.Context) {
	var req struct {
		Value string `json:"value" binding:"required,oneof=light dark system"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid theme value", err)
		return
	}

	theme, err := h.themeService.Switch(req.Value)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to switch theme", err)
		return
	}

	response.Success(c, map[string]interface{}{
		"id":      theme.ID,
		"name":    theme.Name,
		"value":   theme.Value,
		"active":  theme.Active,
		"default": theme.Default,
	})
}