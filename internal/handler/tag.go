package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/api-template-f78c-28HelenNelson/internal/service"
	"github.com/api-template-f78c-28HelenNelson/pkg/response"
)

// TagHandler 封装标签相关 HTTP 处理逻辑
type TagHandler struct {
	tagService *service.TagService
}

// NewTagHandler 创建新的标签处理器
func NewTagHandler(tagService *service.TagService) *TagHandler {
	return &TagHandler{tagService: tagService}
}

// @Summary 创建新标签
// @Description 创建一个新标签，名称唯一
// @Tags Tags
// @Accept json
// @Produce json
// @Param tag body service.CreateTagInput true "标签信息"
// @Success 201 {object} response.Response{data=service.TagOutput}
// @Failure 400 {object} response.Response
// @Failure 409 {object} response.Response
// @Router /api/v1/tags [post]
func (h *TagHandler) Create(c *gin.Context) {
	var input service.CreateTagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, response.ErrInvalidParam, err.Error())
		return
	}

	output, err := h.tagService.Create(c.Request.Context(), input)
	if err != nil {
		response.Error(c, response.ErrInternal, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, output)
}

// @Summary 获取所有标签（支持分页）
// @Description 分页获取标签列表，按创建时间倒序
// @Tags Tags
// @Accept json
// @Produce json
// @Param page query int false "页码（从1开始）" default(1)
// @Param limit query int false "每页数量" default(10)
// @Success 200 {object} response.Response{data=[]service.TagOutput}
// @Failure 400 {object} response.Response
// @Router /api/v1/tags [get]
func (h *TagHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	tags, total, err := h.tagService.List(c.Request.Context(), page, limit)
	if err != nil {
		response.Error(c, response.ErrInternal, err.Error())
		return
	}

	// 构造带分页元信息的响应
	res := map[string]interface{}{
		"list":  tags,
		"total": total,
		"page":  page,
		"limit": limit,
	}
	response.Success(c, http.StatusOK, res)
}

// @Summary 获取单个标签
// @Description 根据 ID 获取标签详情
// @Tags Tags
// @Accept json
// @Produce json
// @Param id path string true "标签ID"
// @Success 200 {object} response.Response{data=service.TagOutput}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /api/v1/tags/{id} [get]
func (h *TagHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, response.ErrInvalidParam, "id is required")
		return
	}

	tag, err := h.tagService.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == service.ErrTagNotFound {
			response.Error(c, response.ErrNotFound, "tag not found")
			return
		}
		response.Error(c, response.ErrInternal, err.Error())
		return
	}

	response.Success(c, http.StatusOK, tag)
}

// @Summary 更新标签
// @Description 更新指定 ID 的标签名称（仅名称可更新）
// @Tags Tags
// @Accept json
// @Produce json
// @Param id path string true "标签ID"
// @Param tag body service.UpdateTagInput true "更新字段"
// @Success 200 {object} response.Response{data=service.TagOutput}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 409 {object} response.Response
// @Router /api/v1/tags/{id} [put]
func (h *TagHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, response.ErrInvalidParam, "id is required")
		return
	}

	var input service.UpdateTagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, response.ErrInvalidParam, err.Error())
		return
	}

	tag, err := h.tagService.Update(c.Request.Context(), id, input)
	if err != nil {
		if err == service.ErrTagNotFound {
			response.Error(c, response.ErrNotFound, "tag not found")
			return
		}
		if err == service.ErrTagDuplicatedName {
			response.Error(c, response.ErrConflict, "tag name already exists")
			return
		}
		response.Error(c, response.ErrInternal, err.Error())
		return
	}

	response.Success(c, http.StatusOK, tag)
}

// @Summary 删除标签
// @Description 根据 ID 删除标签（软删除预留字段，当前为物理删除）
// @Tags Tags
// @Accept json
// @Produce json
// @Param id path string true "标签ID"
// @Success 204 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /api/v1/tags/{id} [delete]
func (h *TagHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, response.ErrInvalidParam, "id is required")
		return
	}

	err := h.tagService.Delete(c.Request.Context(), id)
	if err != nil {
		if err == service.ErrTagNotFound {
			response.Error(c, response.ErrNotFound, "tag not found")
			return
		}
		response.Error(c, response.ErrInternal, err.Error())
		return
	}

	response.NoContent(c)
}