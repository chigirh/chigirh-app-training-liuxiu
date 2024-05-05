package ports

import (
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

// receiving archive
type IArchivementInputPort interface {
	GetArchivement(ctx context.Context, userId models.UserId, chapterId models.ChapterId) (*models.Archivement, error)
	UpdateArchivement(ctx context.Context, archivement models.Archivement) error
}

// repository
type IArchivementRepository interface {
	FetchByUserIdAndThemeId(ctx context.Context, userId models.UserId, themeId models.ThemeId) ([]*models.Archivement, error)
	FetchByUserIdAndChapterId(ctx context.Context, userId models.UserId, chapterId models.ChapterId) (*models.Archivement, error)
	UpdateByArchivementId(ctx context.Context, archivement models.Archivement) error
}
