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
			ChapterId:       string(ch.Id),
			MainExecuteCode: string(ch.MainExecute),
			InitCode:        string(ch.Init),
			Expected:        ch.Expected,
			AnswerCode:      string(ch.Answer),
			Level:           int(ch.Level),
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
			Id:          models.ChapterId(req.Chapter.ChapterId),
			MainExecute: models.Code(req.Chapter.MainExecuteCode),
			Init:        models.Code(req.Chapter.InitCode),
			Expected:    req.Chapter.Expected,
			Answer:      models.Code(req.Chapter.AnswerCode),
			Level:       models.Level(req.Chapter.Level),
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
		ChapterId       string `json:"chapter_id" validate:"required"`
		MainExecuteCode string `json:"main_execute_code" validate:"required"`
		InitCode        string `json:"init_code" validate:"required"`
		Expected        string `json:"expected" validate:"required"`
		AnswerCode      string `json:"answer_code"`
		Level           int    `json:"level" validate:"required"`
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
