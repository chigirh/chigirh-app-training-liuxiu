package interactors

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/errors"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

type ChapterInteractor struct {
	Repository ports.IChapterRepository
}

func (it *ChapterInteractor) AddChapter(ctx context.Context, chapter models.Chapter) error {

	err := it.Repository.Create(ctx, chapter)

	if err != nil {
		return err
	}

	return nil
}

func (it *ChapterInteractor) GetChapter(ctx context.Context, id models.ChapterId) (*models.Chapter, error) {

	chapter, err := it.Repository.FetchBy(ctx, id)

	if err != nil {
		return nil, err
	}

	if chapter == nil {
		return nil, &errors.NotFoundError{Sources: string(id)}
	}

	return chapter, nil
}

func (it *ChapterInteractor) GetAll(ctx context.Context) ([]*models.Chapter, error) {

	chapters, err := it.Repository.FetchAll(ctx)

	if err != nil {
		return nil, err
	}

	return chapters, nil
}

func (it *ChapterInteractor) GetChapterByThemeId(ctx context.Context, id models.ThemeId) ([]*models.Chapter, error) {

	chapters, err := it.Repository.FetchByThemeId(ctx, id)

	if err != nil {
		return nil, err
	}

	return chapters, nil
}

// di
func NewIChapterInputPort(repository ports.IChapterRepository) ports.IChapterInputPort {
	return &ChapterInteractor{
		Repository: repository,
	}
}
