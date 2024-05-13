package ports

import (
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

// receiving theme
type IThemeInputPort interface {
	GetAll(ctx context.Context) ([]*models.Theme, error)
	GetTheme(ctx context.Context, userId models.UserId, themeId models.ThemeId) (*models.Theme, error)
	AddTheme(ctx context.Context, theme models.Theme) error
	UpdateChapter(ctx context.Context, theme models.Theme, chapters []models.ThemeChapters) error
}

// repository
type IThemeRepository interface {
	FetchAll(ctx context.Context) ([]*models.Theme, error)
	FetchByThemeId(ctx context.Context, themeId models.ThemeId) (*models.Theme, error)
	Create(ctx context.Context, theme models.Theme) error
	UpdateChapter(ctx context.Context, theme models.Theme, chapters []models.ThemeChapters) error
}
