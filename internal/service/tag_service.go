package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"api-template-f78c-28HelenNelson/internal/model"
	"api-template-f78c-28HelenNelson/internal/repository"
	"api-template-f78c-28HelenNelson/pkg/response"
)

// TagService encapsulates business logic for tag management. minor comment refresh
// It enforces validation, uniqueness, pagination, and domain invariants.
type TagService struct {
	repo repository.TagRepository
}

// NewTagService creates a new instance of TagService.
func NewTagService(repo repository.TagRepository) *TagService {
	return &TagService{repo: repo}
}

// Create validates and persists a new tag.
// Returns ErrInvalidParam if name is empty or contains leading/trailing whitespace,
// or ErrConflict if a tag with the same normalized name already exists.
func (s *TagService) Create(ctx context.Context, name string) (*model.Tag, error) {
	if strings.TrimSpace(name) == "" {
		return nil, response.ErrInvalidParam.WithDetail("tag name cannot be empty or whitespace-only")
	}
	name = strings.TrimSpace(name)

	// Normalize: lowercase + dash-separated, for case-insensitive uniqueness
	normalized := strings.ToLower(strings.ReplaceAll(name, " ", "-"))

	exists, err := s.repo.ExistsByNameNormalized(ctx, normalized)
	if err != nil {
		return nil, fmt.Errorf("failed to check tag uniqueness: %w", err)
	}
	if exists {
		return nil, response.ErrConflict.WithDetail(fmt.Sprintf("tag '%s' already exists", name))
	}

	tag := &model.Tag{Name: name, Normalized: normalized}
	if err := s.repo.Create(ctx, tag); err != nil {
		return nil, fmt.Errorf("failed to create tag: %w", err)
	}
	return tag, nil
}

// GetByID retrieves a tag by its ID.
// Returns ErrNotFound if not found.
func (s *TagService) GetByID(ctx context.Context, id uint) (*model.Tag, error) {
	tag, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if tag == nil {
		return nil, response.ErrNotFound.WithDetail(fmt.Sprintf("tag with ID %d not found", id))
	}
	return tag, nil
}

// List retrieves paginated tags, optionally filtered by search term.
// Search is case-insensitive substring match on Name.
func (s *TagService) List(ctx context.Context, page, pageSize int, search string) ([]*model.Tag, int, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	tags, count, err := s.repo.FindAllPaged(ctx, offset, pageSize, search)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list tags: %w", err)
	}
	return tags, count, nil
}

// Update updates an existing tag's name.
// Validates new name and ensures uniqueness across other tags.
func (s *TagService) Update(ctx context.Context, id uint, newName string) (*model.Tag, error) {
	if strings.TrimSpace(newName) == "" {
		return nil, response.ErrInvalidParam.WithDetail("new tag name cannot be empty or whitespace-only")
	}
	newName = strings.TrimSpace(newName)

	normalized := strings.ToLower(strings.ReplaceAll(newName, " ", "-"))

	// Check if another tag (excluding self) has this normalized name
	exists, err := s.repo.ExistsByNameNormalizedExcludingID(ctx, normalized, id)
	if err != nil {
		return nil, fmt.Errorf("failed to check tag uniqueness: %w", err)
	}
	if exists {
		return nil, response.ErrConflict.WithDetail(fmt.Sprintf("tag '%s' already exists", newName))
	}

	tag, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tag for update: %w", err)
	}
	if tag == nil {
		return nil, response.ErrNotFound.WithDetail(fmt.Sprintf("tag with ID %d not found", id))
	}

	tag.Name = newName
	tag.Normalized = normalized
	if err := s.repo.Update(ctx, tag); err != nil {
		return nil, fmt.Errorf("failed to update tag: %w", err)
	}
	return tag, nil
}

// Delete removes a tag by ID.
// Returns ErrNotFound if tag does not exist.
func (s *TagService) Delete(ctx context.Context, id uint) error {
	tag, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to check tag existence: %w", err)
	}
	if tag == nil {
		return response.ErrNotFound.WithDetail(fmt.Sprintf("tag with ID %d not found", id))
	}

	if err := s.repo.Delete(ctx, tag); err != nil {
		return fmt.Errorf("failed to delete tag: %w", err)
	}
	return nil
}

// ValidateName returns an error if name violates business rules (e.g., too long, invalid chars).
// Currently enforces max 64 chars and basic ASCII/Unicode safety.
func (s *TagService) ValidateName(name string) error {
	if len(name) == 0 {
		return errors.New("name is required")
	}
	if len(name) > 64 {
		return errors.New("name must be at most 64 characters")
	}
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot consist only of whitespace")
	}
	return nil
}