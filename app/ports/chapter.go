package ports

import (
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

// ports
type IChapterInputPort interface {
	AddChapter(ctx context.Context, chapter models.Chapter) error
	GetChapter(ctx context.Context, id models.ChapterId) (*models.Chapter, error)
}

// repositories
type IChapterRepository interface {
	Create(ctx context.Context, chapter models.Chapter) error
	FetchBy(ctx context.Context, id models.ChapterId) (*models.Chapter, error)
}
