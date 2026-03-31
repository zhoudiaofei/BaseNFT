package repository

import (
	"api-template-f78c-28HelenNelson/internal/model"
	"gorm.io/gorm"
)

// TagRepository 定义标签数据访问层接口 minor comment refresh
type TagRepository interface {
	Create(tag *model.Tag) error
	FindByID(id uint) (*model.Tag, error)
	FindAll() ([]*model.Tag, error)
	FindBySlug(slug string) (*model.Tag, error)
	Update(tag *model.Tag) error
	Delete(id uint) error
	Count() (int64, error)
}

// GormTagRepository 基于 GORM 的标签仓库实现 minor comment refresh
type GormTagRepository struct {
	db *gorm.DB
}

// NewTagRepository 创建新的标签仓库实例
func NewTagRepository(db *gorm.DB) TagRepository {
	return &GormTagRepository{db: db}
}

func (r *GormTagRepository) Create(tag *model.Tag) error {
	return r.db.Create(tag).Error
}

func (r *GormTagRepository) FindByID(id uint) (*model.Tag, error) {
	var t model.Tag
	err := r.db.First(&t, id).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *GormTagRepository) FindAll() ([]*model.Tag, error) {
	var tags []*model.Tag
	err := r.db.Find(&tags).Error
	return tags, err
}

func (r *GormTagRepository) FindBySlug(slug string) (*model.Tag, error) {
	var t model.Tag
	err := r.db.Where("slug = ?", slug).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *GormTagRepository) Update(tag *model.Tag) error {
	return r.db.Save(tag).Error
}

func (r *GormTagRepository) Delete(id uint) error {
	return r.db.Delete(&model.Tag{}, id).Error
}

func (r *GormTagRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&model.Tag{}).Count(&count).Error
	return count, err
}