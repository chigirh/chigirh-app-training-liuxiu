package ports

import (
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

// repository
type IRevisionRepository interface {
	FetchByArchivementId(ctx context.Context, archivementId models.ArchivementId) (*models.Revision, error)
	FetchByArchivementIdAndVersion(ctx context.Context, archivementId models.ArchivementId, version models.RevisionVersion) (*models.Revision, error)
	Create(ctx context.Context, revision models.Revision) error
	UpdateByChapterIdAndVersion(ctx context.Context, revision models.Revision) error
}
