package interactors

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

type ArchivementInteractor struct {
	ArchivementRepository ports.IArchivementRepository
	RevisionRepository    ports.IRevisionRepository
}

func (it *ArchivementInteractor) GetArchivement(
	ctx context.Context,
	userId models.UserId,
	chapterId models.ChapterId,
) (*models.Archivement, error) {

	archivement, err := it.ArchivementRepository.FetchByUserIdAndChapterId(ctx, userId, chapterId)

	if err != nil {
		return nil, err
	}

	if archivement == nil {
		id := models.NewArchivementId("")
		archivement := &models.Archivement{
			ArchivementId: id,
			ChapterId:     chapterId,
			UserId:        userId,
			Order:         0,
			Status:        models.NewArchivementStatus("1"),
		}
		err := it.ArchivementRepository.UpdateByArchivementId(ctx, *archivement)

		if err != nil {
			return nil, err
		}
	}

	revision, err := it.RevisionRepository.FetchByArchivementId(ctx, archivement.ArchivementId)

	if err != nil {
		return nil, err
	}

	if revision == nil {
		revision = &models.Revision{
			ArchivementId:  archivement.ArchivementId,
			Version:        models.NewRevisionVersion(0),
			Status:         models.NewArchivementStatus("1"),
			Code:           models.Code(""),
			Comment:        "",
			Result:         "",
			IsCompileError: false,
		}
	}

	archivement.Revision = *revision

	return archivement, nil
}

func (it *ArchivementInteractor) UpdateArchivement(ctx context.Context, archivement models.Archivement) error {

	revision := archivement.Revision
	archivement.Status = revision.Status

	it.ArchivementRepository.UpdateByArchivementId(ctx, archivement)

	r, err := it.RevisionRepository.FetchByArchivementIdAndVersion(ctx, revision.ArchivementId, revision.Version)

	if err != nil {
		return err
	}

	if r != nil {
		err = it.RevisionRepository.UpdateByChapterIdAndVersion(ctx, revision)
	} else {
		err = it.RevisionRepository.Create(ctx, revision)
	}

	if err != nil {
		return err
	}

	return nil

}

// di
func NewArchivementInputPort(
	archivementRepository ports.IArchivementRepository,
	revisionRepository ports.IRevisionRepository,
) ports.IArchivementInputPort {
	return &ArchivementInteractor{
		ArchivementRepository: archivementRepository,
		RevisionRepository:    revisionRepository,
	}
}
