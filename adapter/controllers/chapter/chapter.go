package chapter

import (
	"context"
	"net/http"

	"chigirh-app-trainning-liuxiu/adapter/controllers"
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/conf/config"
	"chigirh-app-trainning-liuxiu/domain/models"

	"github.com/labstack/echo"
)

type IChapterApi interface {
	Get(ctx context.Context) func(c echo.Context) error
	Post(ctx context.Context) func(c echo.Context) error
	AllGet(ctx context.Context) func(c echo.Context) error
	ListGet(ctx context.Context) func(c echo.Context) error
}

type ChapterController struct {
	RequestMapper controllers.RequestMapper
	StudyAuthPort ports.IStudyAuthPort
	InputPort     ports.IChapterInputPort
}

func (it *ChapterController) Get(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {
		sk, err := it.RequestMapper.GetSessionKey(c)
		if err != nil {
			return err
		}

		_, err = it.StudyAuthPort.GetAuthorizedUser(ctx, models.SessionKey(sk))

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		id := c.Param("chapterId")

		ch, err := it.InputPort.GetChapter(ctx, models.ChapterId(id))

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		res := new(GetResponse)
		res.Chapter = ChapterData{
			ChapterId:        string(ch.Id),
			MainCode:         string(ch.Main),
			ExampleCode:      string(ch.Example),
			Expected:         ch.Expected,
			BestPracticeCode: string(ch.BestPractice),
			Level:            int(ch.Level),
			Exercise:         ch.Exercise,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (it *ChapterController) Post(ctx context.Context) func(c echo.Context) error {

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

		ch := models.Chapter{
			Id:           models.ChapterId(req.Chapter.ChapterId),
			Main:         models.Code(req.Chapter.MainCode),
			Example:      models.Code(req.Chapter.ExampleCode),
			Expected:     req.Chapter.Expected,
			BestPractice: models.Code(req.Chapter.BestPracticeCode),
			Level:        models.Level(req.Chapter.Level),
			Exercise:     req.Chapter.Exercise,
		}

		it.InputPort.AddChapter(ctx, ch)

		return c.JSON(http.StatusOK, controllers.DefaultResponse)
	}
}

func (it *ChapterController) AllGet(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {
		mk, err := it.RequestMapper.GetMasterKey(c)
		if err != nil {
			return err
		}

		if config.Server.MasterKey != string(mk) {
			return c.JSON(http.StatusForbidden, controllers.DefaultResponse)
		}

		chapters, err := it.InputPort.GetAll(ctx)

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		dtos := []ChapterData{}
		for _, chapter := range chapters {
			dto := ChapterData{
				ChapterId:        string(chapter.Id),
				MainCode:         string(chapter.Main),
				ExampleCode:      string(chapter.Example),
				Expected:         chapter.Expected,
				BestPracticeCode: string(chapter.BestPractice),
				Level:            int(chapter.Level),
				Exercise:         chapter.Exercise,
			}
			dtos = append(dtos, dto)
		}
		res := new(AllGetResponse)
		res.Chapters = dtos

		return c.JSON(http.StatusOK, res)
	}
}

func (it *ChapterController) ListGet(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {
		mk, err := it.RequestMapper.GetMasterKey(c)
		if err != nil {
			return err
		}

		if config.Server.MasterKey != string(mk) {
			return c.JSON(http.StatusForbidden, controllers.DefaultResponse)
		}

		id := c.Param("themeId")
		chapters, err := it.InputPort.GetChapterByThemeId(ctx, models.ThemeId(id))

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		dtos := []ChapterData{}
		for _, chapter := range chapters {
			dto := ChapterData{
				ChapterId:        string(chapter.Id),
				MainCode:         string(chapter.Main),
				ExampleCode:      string(chapter.Example),
				Expected:         chapter.Expected,
				BestPracticeCode: string(chapter.BestPractice),
				Level:            int(chapter.Level),
				Exercise:         chapter.Exercise,
			}
			dtos = append(dtos, dto)
		}
		res := new(AllGetResponse)
		res.Chapters = dtos

		return c.JSON(http.StatusOK, res)
	}
}

// dto -->
type (
	GetResponse struct {
		Chapter ChapterData `json:"chapter" validate:"required"`
	}

	PostRequest struct {
		Chapter ChapterData `json:"chapter" validate:"required"`
	}

	AllGetResponse struct {
		Chapters []ChapterData `json:"chapters" validate:"required"`
	}

	ListGetResponse struct {
		Chapters []ChapterData `json:"chapters" validate:"required"`
	}

	ChapterData struct {
		ChapterId        string `json:"chapter_id" validate:"required"`
		MainCode         string `json:"main_code" validate:"required"`
		ExampleCode      string `json:"example_code" validate:"required"`
		Expected         string `json:"expected" validate:"required"`
		BestPracticeCode string `json:"best_practice_code"`
		Level            int    `json:"level" validate:"required"`
		Exercise         string `json:"exercise"`
	}
)

func NewChapterController(
	RequestMapper controllers.RequestMapper,
	StudyAuthPort ports.IStudyAuthPort,
	InputPort ports.IChapterInputPort,
) IChapterApi {
	return &ChapterController{
		RequestMapper: RequestMapper,
		StudyAuthPort: StudyAuthPort,
		InputPort:     InputPort,
	}
}
