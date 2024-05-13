package ports

import (
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

// receiving chapter
type IChapterInputPort interface {
	AddChapter(ctx context.Context, chapter models.Chapter) error
	GetChapter(ctx context.Context, id models.ChapterId) (*models.Chapter, error)
	GetAll(ctx context.Context) ([]*models.Chapter, error)
	GetChapterByThemeId(ctx context.Context, id models.ThemeId) ([]*models.Chapter, error)
}

// repositories
type IChapterRepository interface {
	Create(ctx context.Context, chapter models.Chapter) error
	FetchBy(ctx context.Context, id models.ChapterId) (*models.Chapter, error)
	FetchAll(ctx context.Context) ([]*models.Chapter, error)
	FetchByThemeId(ctx context.Context, id models.ThemeId) ([]*models.Chapter, error)
}
