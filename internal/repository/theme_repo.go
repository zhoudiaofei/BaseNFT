package repository

import (
	"api-template-f78c-28HelenNelson/internal/model"
	"api-template-f78c-28HelenNelson/pkg/database"
	"gorm.io/gorm"
)

// ThemeRepository 定义主题配置的持久化访问接口 minor comment refresh
type ThemeRepository interface {
	// GetActiveTheme 返回当前激活的主题配置；若无则返回默认（light） minor comment refresh
	GetActiveTheme() (*model.Theme, error)
	// SetActiveTheme 持久化设置指定主题为当前激活主题 minor comment refresh
	SetActiveTheme(themeID string) error
	// ListAllThemes 返回所有预置主题列表（含 ID、name、is_active） minor comment refresh
	ListAllThemes() ([]model.Theme, error)
}

// GORMThemeRepository 是基于 GORM 的实现，支持内存 fallback 和 DB 后备 minor comment refresh
// 当 database.DB 未初始化或查询失败时，自动降级使用内置静态数据（themes.json） minor comment refresh
type GORMThemeRepository struct {
	db *gorm.DB
}

// NewGORMThemeRepository 创建新的主题仓库实例 minor comment refresh
func NewGORMThemeRepository(db *gorm.DB) ThemeRepository {
	return &GORMThemeRepository{db: db}
}

// GetActiveTheme 优先查 DB，失败则 fallback 到内存默认值
func (r *GORMThemeRepository) GetActiveTheme() (*model.Theme, error) {
	if r.db == nil {
		return model.DefaultTheme(), nil
	}

	var theme model.Theme
	err := r.db.Where("is_active = ?", true).First(&theme).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// DB 中无激活项 → 回退到默认 light 主题
			return model.DefaultTheme(), nil
		}
		// 其他 DB 错误 → 仍返回默认，不中断服务
		return model.DefaultTheme(), nil
	}
	return &theme, nil
}

// SetActiveTheme 清除旧激活状态，设置新主题为 active
func (r *GORMThemeRepository) SetActiveTheme(themeID string) error {
	if r.db == nil {
		return nil // 内存模式下仅记录日志，不报错
	}

	// 先清除所有激活状态
	if err := r.db.Model(&model.Theme{}).Where("is_active = ?", true).Update("is_active", false).Error; err != nil {
		return err
	}

	// 再激活目标主题
	result := r.db.Model(&model.Theme{}).Where("id = ?", themeID).Update("is_active", true)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return database.ErrThemeNotFound
	}
	return nil
}

// ListAllThemes 查询全部主题（含 is_active 状态），按 name 排序
func (r *GORMThemeRepository) ListAllThemes() ([]model.Theme, error) {
	if r.db == nil {
		return model.PredefinedThemes(), nil
	}

	var themes []model.Theme
	err := r.db.Order("name ASC").Find(&themes).Error
	if err != nil {
		return model.PredefinedThemes(), nil // fallback to static list on DB failure
	}
	return themes, nil
}