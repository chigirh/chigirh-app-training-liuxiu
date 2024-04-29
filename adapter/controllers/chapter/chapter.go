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
}

type ChapterController struct {
	requestMapper controllers.RequestMapper
	inputPort     ports.IChapterInputPort
}

func (it *ChapterController) Get(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {
		id := c.Param("chapterId")

		mk, err := it.requestMapper.GetMasterKey(c)
		if err != nil {
			return err
		}

		if config.Server.MasterKey != string(mk) {
			return c.JSON(http.StatusForbidden, controllers.DefaultResponse)
		}

		ch, err := it.inputPort.GetChapter(ctx, models.ChapterId(id))

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
		mk, err := it.requestMapper.GetMasterKey(c)
		if err != nil {
			return err
		}

		if config.Server.MasterKey != string(mk) {
			return c.JSON(http.StatusForbidden, controllers.DefaultResponse)
		}

		req := new(PostRequest)

		if err := it.requestMapper.Parse(c, req); err != nil {
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

		it.inputPort.AddChapter(ctx, ch)

		return c.JSON(http.StatusOK, controllers.DefaultResponse)
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
	requestMapper controllers.RequestMapper,
	inputPost ports.IChapterInputPort,
) IChapterApi {
	return &ChapterController{
		requestMapper: requestMapper,
		inputPort:     inputPost,
	}
}
