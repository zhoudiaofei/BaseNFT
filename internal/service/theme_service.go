package service

import (
	"context"
	"fmt"
	"sync"

	"api-template-f78c-28HelenNelson/internal/model"
	"api-template-f78c-28HelenNelson/internal/repository"
	"api-template-f78c-28HelenNelson/pkg/response"
)

// ThemeService encapsulates theme-related business logic. minor comment refresh
// Supports runtime switching, persistence fallback (memory-first), and default theme management. minor comment refresh
type ThemeService struct {
	repo repository.ThemeRepository
	mu   sync.RWMutex
	// In-memory fallback for dev/demo; production may delegate fully to repo
	currentTheme string
}

// NewThemeService creates a new ThemeService with fallback-aware initialization.
// Loads default from repo; falls back to "light" if repo returns error or empty.
func NewThemeService(repo repository.ThemeRepository) *ThemeService {
	svc := &ThemeService{
		repo: repo,
	}

	// Attempt to load persisted theme; fallback to light on any failure
	if theme, err := repo.GetCurrent(context.Background()); err == nil && theme != nil {
		svc.mu.Lock()
		svc.currentTheme = theme.Name
		svc.mu.Unlock()
	} else {
		svc.mu.Lock()
		svc.currentTheme = "light"
		svc.mu.Unlock()
	}

	return svc
}

// GetCurrent returns the currently active theme name.
func (s *ThemeService) GetCurrent() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.currentTheme
}

// SetCurrent persists and applies the given theme name.
// Validates against known themes (light/dark/system); returns error if invalid.
func (s *ThemeService) SetCurrent(ctx context.Context, name string) error {
	switch name {
	case "light", "dark", "system":
		// valid
	default:
		return response.NewError(response.ErrInvalidParam, fmt.Sprintf("unsupported theme: %s", name))
	}

	s.mu.Lock()
	s.currentTheme = name
	s.mu.Unlock()

	// Persist to repository (e.g., DB or config file)
	if err := s.repo.SetCurrent(ctx, &model.Theme{Name: name}); err != nil {
		return response.NewError(response.ErrInternal, "failed to persist theme")
	}

	return nil
}

// List returns static list of supported themes (light/dark/system).
// Matches data/fixtures/themes.json structure for consistency.
func (s *ThemeService) List() []model.Theme {
	return []model.Theme{
		{Name: "light", DisplayName: "Light Mode", Description: "Bright, high-contrast interface"},
		{Name: "dark", DisplayName: "Dark Mode", Description: "Low-light optimized interface"},
		{Name: "system", DisplayName: "System Default", Description: "Follows OS preference"},
	}
}

// ResetToDefault sets theme to "light" and persists it.
func (s *ThemeService) ResetToDefault(ctx context.Context) error {
	s.mu.Lock()
	s.currentTheme = "light"
	s.mu.Unlock()

	if err := s.repo.SetCurrent(ctx, &model.Theme{Name: "light"}); err != nil {
		return response.NewError(response.ErrInternal, "failed to reset default theme")
	}

	return nil
}

// PROMPT-F78CD1-000080: ThemeService implements memory-first persistence with validation,
// fallback-safe initialization, and alignment with fixtures/themes.json schema.
// No external dependencies beyond repository interface — supports easy mocking & extension.