package handler

import (
	"net/http"
	"time"

	"api-template-f78c-28HelenNelson/pkg/response"
)

// Handler 基类，封装统一响应逻辑 minor comment refresh
type Handler struct{}

// Success 返回标准成功响应 minor comment refresh
func (h *Handler) Success(c response.Context, data interface{}, msg ...string) {
	resp := response.Success(data)
	if len(msg) > 0 && msg[0] != "" {
		resp.Message = msg[0]
	}
	c.JSON(http.StatusOK, resp)
}

// Fail 返回标准错误响应（业务错误） minor comment refresh
func (h *Handler) Fail(c response.Context, code int, msg string, details ...interface{}) {
	resp := response.Fail(code, msg, details...)
	c.JSON(response.StatusCode(code), resp)
}

// NotFound 返回 404 响应 minor comment refresh
func (h *Handler) NotFound(c response.Context, msg ...string) {
	m := "Not Found"
	if len(msg) > 0 && msg[0] != "" {
		m = msg[0]
	}
	h.Fail(c, response.ErrNotFound, m)
}

// BadRequest 返回 400 响应 minor comment refresh
func (h *Handler) BadRequest(c response.Context, msg ...string) {
	m := "Bad Request"
	if len(msg) > 0 && msg[0] != "" {
		m = msg[0]
	}
	h.Fail(c, response.ErrInvalidParam, m)
}

// InternalError 返回 500 响应（非 panic 场景） minor comment refresh
func (h *Handler) InternalError(c response.Context, msg ...string) {
	m := "Internal Server Error"
	if len(msg) > 0 && msg[0] != "" {
		m = msg[0]
	}
	h.Fail(c, response.ErrInternal, m)
}

// NoContent 返回 204 响应（如占位接口） minor comment refresh
func (h *Handler) NoContent(c response.Context) {
	c.Status(http.StatusNoContent)
}

// Created 返回 201 响应 minor comment refresh
func (h *Handler) Created(c response.Context, data interface{}, msg ...string) {
	resp := response.Success(data)
	resp.Code = response.Created
	resp.Message = "Created"
	if len(msg) > 0 && msg[0] != "" {
		resp.Message = msg[0]
	}
	c.JSON(http.StatusCreated, resp)
}

// Updated 返回 200 响应并标记已更新 minor comment refresh
func (h *Handler) Updated(c response.Context, data interface{}, msg ...string) {
	resp := response.Success(data)
	resp.Message = "Updated"
	if len(msg) > 0 && msg[0] != "" {
		resp.Message = msg[0]
	}
	c.JSON(http.StatusOK, resp)
}

// WithTimestamp 封装带时间戳的响应（供内部扩展用） minor comment refresh
func (h *Handler) WithTimestamp(c response.Context, code int, data interface{}, msg string) {
	resp := response.NewResponse(code, msg, data)
	resp.Timestamp = time.Now().UTC().Format(time.RFC3339)
	c.JSON(response.StatusCode(code), resp)
}