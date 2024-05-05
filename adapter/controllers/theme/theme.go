package theme

import (
	"context"
	"net/http"

	"chigirh-app-trainning-liuxiu/adapter/controllers"
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/models"

	"github.com/labstack/echo"
)

type IThemeApi interface {
	Get(ctx context.Context) func(c echo.Context) error
}

type ThemeController struct {
	RequestMapper controllers.RequestMapper
	StudyAuthPort ports.IStudyAuthPort
	InputPort     ports.IThemeInputPort
}

func (it *ThemeController) Get(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {

		sk, err := it.RequestMapper.GetSessionKey(c)
		if err != nil {
			return err
		}

		user, err := it.StudyAuthPort.GetAuthorizedUser(ctx, models.SessionKey(sk))

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		theme, err := it.InputPort.GetTheme(ctx, user.UserId, user.ThemeId)

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		res := new(GetResponse)

		archivementDtos := []ArchivementDto{}

		for _, e := range theme.Archivements {
			dto := ArchivementDto{
				ArchivementId: string(e.ArchivementId),
				ChapterId:     string(e.ChapterId),
				Order:         e.Order,
				Status:        string(e.Status),
			}
			archivementDtos = append(archivementDtos, dto)
		}

		res.Theme = ThemeDto{
			ThemeId:      string(theme.ThemeId),
			Theme:        string(theme.Theme),
			Description:  theme.Description,
			Archivements: archivementDtos,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// dto -->
type (
	GetResponse struct {
		Theme ThemeDto `json:"theme" validate:"required"`
	}

	ThemeDto struct {
		ThemeId      string           `json:"theme_id" validate:"required,max=64"`
		Theme        string           `json:"theme" validate:"required"`
		Description  string           `json:"description" validate:"required"`
		Archivements []ArchivementDto `json:"archivements"`
	}

	ArchivementDto struct {
		ArchivementId string `json:"archivement_id" validate:"required,max=64"`
		ChapterId     string `json:"chapter_id" validate:"required,max=64"`
		Order         int    `json:"order" validate:"required"`
		Status        string `json:"status" validate:"required"`
	}
)

func NewThemeController(
	RequestMapper controllers.RequestMapper,
	StudyAuthPort ports.IStudyAuthPort,
	InputPort ports.IThemeInputPort,
) IThemeApi {
	return &ThemeController{
		RequestMapper: RequestMapper,
		StudyAuthPort: StudyAuthPort,
		InputPort:     InputPort,
	}
}
