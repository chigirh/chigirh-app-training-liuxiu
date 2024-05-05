package ports

import (
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

// receiving theme
type IThemeInputPort interface {
	GetTheme(ctx context.Context, userId models.UserId, themeId models.ThemeId) (*models.Theme, error)
}

// repository
type IThemeRepository interface {
	FetchByThemeId(ctx context.Context, themeId models.ThemeId) (*models.Theme, error)
}
