package interactors

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/errors"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

type ThemeInteractor struct {
	ThemeRepository   ports.IThemeRepository
	ArchiveRepository ports.IArchivementRepository
}

func (it *ThemeInteractor) GetAll(ctx context.Context) ([]*models.Theme, error) {

	themes, err := it.ThemeRepository.FetchAll(ctx)

	if err != nil {
		return nil, err
	}

	return themes, nil
}

func (it *ThemeInteractor) GetTheme(ctx context.Context, userId models.UserId, themeId models.ThemeId) (*models.Theme, error) {

	theme, err := it.ThemeRepository.FetchByThemeId(ctx, themeId)

	if err != nil {
		return nil, err
	}

	if theme == nil {
		return nil, &errors.NotFoundError{Sources: string(themeId)}
	}

	archivements, err := it.ArchiveRepository.FetchByUserIdAndThemeId(ctx, userId, themeId)

	if err != nil {
		return nil, err
	}

	theme.Archivements = models.Archivements(archivements)

	return theme, nil
}

func (it *ThemeInteractor) AddTheme(ctx context.Context, theme models.Theme) error {

	t, err := it.ThemeRepository.FetchByThemeId(ctx, theme.ThemeId)

	if err != nil {
		return err
	}

	if t != nil {
		return &errors.AlreadyExistsError{Sources: string(theme.ThemeId)}
	}

	err = it.ThemeRepository.Create(ctx, theme)

	if err != nil {
		return err
	}

	return nil
}

func (it *ThemeInteractor) UpdateChapter(ctx context.Context, theme models.Theme, chapters []models.ThemeChapters) error {

	t, err := it.ThemeRepository.FetchByThemeId(ctx, theme.ThemeId)

	if err != nil {
		return err
	}

	if t == nil {
		return &errors.NotFoundError{Sources: string(theme.ThemeId)}
	}
	err = it.ThemeRepository.UpdateChapter(ctx, theme, chapters)

	if err != nil {
		return err
	}

	return nil
}

// di
func NewIThemeInputPort(
	ThemeRepository ports.IThemeRepository,
	ArchiveRepository ports.IArchivementRepository,
) ports.IThemeInputPort {
	return &ThemeInteractor{
		ThemeRepository:   ThemeRepository,
		ArchiveRepository: ArchiveRepository,
	}
}
