package theme

import (
	"context"
	"net/http"

	"chigirh-app-trainning-liuxiu/adapter/controllers"
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/conf/config"
	"chigirh-app-trainning-liuxiu/domain/models"

	"github.com/labstack/echo"
)

type IThemeApi interface {
	AllGet(ctx context.Context) func(c echo.Context) error
	Get(ctx context.Context) func(c echo.Context) error
	Post(ctx context.Context) func(c echo.Context) error
	ChapterPut(ctx context.Context) func(c echo.Context) error
}

type ThemeController struct {
	RequestMapper controllers.RequestMapper
	StudyAuthPort ports.IStudyAuthPort
	InputPort     ports.IThemeInputPort
}

func (it *ThemeController) AllGet(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {
		mk, err := it.RequestMapper.GetMasterKey(c)
		if err != nil {
			return err
		}

		if config.Server.MasterKey != string(mk) {
			return c.JSON(http.StatusForbidden, controllers.DefaultResponse)
		}

		themes, err := it.InputPort.GetAll(ctx)

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		dtos := []ThemeDto{}
		for _, theme := range themes {
			dto := ThemeDto{
				ThemeId:     string(theme.ThemeId),
				Theme:       string(theme.Theme),
				Description: theme.Description,
			}
			dtos = append(dtos, dto)
		}
		res := new(AllGetResponse)
		res.Themes = dtos

		return c.JSON(http.StatusOK, res)
	}
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
			ThemeId:     string(theme.ThemeId),
			Theme:       string(theme.Theme),
			Description: theme.Description,
		}
		res.Archivements = archivementDtos

		return c.JSON(http.StatusOK, res)
	}
}

func (it *ThemeController) Post(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {

		mk, err := it.RequestMapper.GetMasterKey(c)
		if err != nil {
			return err
		}

		if config.Server.MasterKey != string(mk) {
			return c.JSON(http.StatusForbidden, controllers.DefaultResponse)
		}

		req := new(PostRequest)

		if err := it.RequestMapper.Parse(c, req); err != nil {
			return err
		}

		theme := models.Theme{
			ThemeId:     models.ThemeId(req.Theme.ThemeId),
			Theme:       req.Theme.Theme,
			Description: req.Theme.Description,
		}

		err = it.InputPort.AddTheme(ctx, theme)
		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		return c.JSON(http.StatusOK, controllers.DefaultResponse)
	}
}

func (it *ThemeController) ChapterPut(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {

		mk, err := it.RequestMapper.GetMasterKey(c)
		if err != nil {
			return err
		}

		if config.Server.MasterKey != string(mk) {
			return c.JSON(http.StatusForbidden, controllers.DefaultResponse)
		}

		req := new(ChapterPutRequest)

		if err := it.RequestMapper.Parse(c, req); err != nil {
			return err
		}

		chapters := []models.ThemeChapters{}

		for _, e := range req.Chapters {
			c := models.ThemeChapters{
				ChapterId: models.ChapterId(e.ChapterId),
				Order:     e.Order,
			}
			chapters = append(chapters, c)
		}

		err = it.InputPort.UpdateChapter(ctx, models.Theme{ThemeId: models.ThemeId(req.ThemeId)}, chapters)
		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		return c.JSON(http.StatusOK, controllers.DefaultResponse)
	}
}

// dto -->
type (
	GetResponse struct {
		Theme        ThemeDto         `json:"theme" validate:"required"`
		Archivements []ArchivementDto `json:"archivements"`
	}

	AllGetResponse struct {
		Themes []ThemeDto `json:"themes"`
	}

	PostRequest struct {
		Theme ThemeDto `json:"theme" validate:"required"`
	}

	ChapterPutRequest struct {
		ThemeId  string            `json:"theme_id" validate:"required,max=64"`
		Chapters []ThemeChapterDto `json:"chapters"`
	}

	ThemeDto struct {
		ThemeId     string `json:"theme_id" validate:"required,max=64"`
		Theme       string `json:"theme" validate:"required"`
		Description string `json:"description" validate:"required"`
	}

	ThemeChapterDto struct {
		ChapterId string `json:"chapter_id" validate:"required"`
		Order     int    `json:"order" validate:"required"`
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
