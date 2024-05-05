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
