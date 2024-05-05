package archivement

import (
	"context"
	"net/http"

	"chigirh-app-trainning-liuxiu/adapter/controllers"
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/models"

	"github.com/labstack/echo"
)

type IArchivementApi interface {
	Get(ctx context.Context) func(c echo.Context) error
	Post(ctx context.Context) func(c echo.Context) error
}

type ArchivementController struct {
	RequestMapper controllers.RequestMapper
	StudyAuthPort ports.IStudyAuthPort
	InputPort     ports.IArchivementInputPort
}

func (it *ArchivementController) Get(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {

		sk, err := it.RequestMapper.GetSessionKey(c)
		if err != nil {
			return err
		}

		user, err := it.StudyAuthPort.GetAuthorizedUser(ctx, models.SessionKey(sk))

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		id := c.Param("chapterId")

		archivement, err := it.InputPort.GetArchivement(ctx, user.UserId, models.ChapterId(id))

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		revision := archivement.Revision

		res := new(GetResponse)

		res.Archivement = ArchivementDto{
			ArchivementId:  string(archivement.ArchivementId),
			Status:         string(revision.Status),
			Version:        int(revision.Version),
			Code:           string(revision.Code),
			Comment:        revision.Comment,
			Result:         revision.Result,
			IsCompileError: revision.IsCompileError,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (it *ArchivementController) Post(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {

		sk, err := it.RequestMapper.GetSessionKey(c)
		if err != nil {
			return err
		}

		_, err = it.StudyAuthPort.GetAuthorizedUser(ctx, models.SessionKey(sk))

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		req := new(PostRequest)
		if err := it.RequestMapper.Parse(c, req); err != nil {
			return err
		}

		archivement := models.Archivement{
			ArchivementId: models.NewArchivementId(req.Archivement.ArchivementId),
			ChapterId:     models.ChapterId(""), //  not use
			Order:         0,                    //  not use
			Status:        models.NewArchivementStatus(req.Archivement.Status),
			Revision: models.Revision{
				ArchivementId:  models.NewArchivementId(req.Archivement.ArchivementId),
				Version:        models.RevisionVersion(req.Archivement.Version),
				Status:         models.NewArchivementStatus(req.Archivement.Status),
				Code:           models.Code(req.Archivement.Code),
				Comment:        req.Archivement.Comment,
				Result:         req.Archivement.Result,
				IsCompileError: req.Archivement.IsCompileError,
			},
		}

		err = it.InputPort.UpdateArchivement(ctx, archivement)

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		res := controllers.DefaultResponse

		return c.JSON(http.StatusOK, res)
	}
}

// dto -->
type (
	GetResponse struct {
		Archivement ArchivementDto `json:"archivement" validate:"required"`
	}

	PostRequest struct {
		Archivement ArchivementDto `json:"archivement" validate:"required"`
	}

	ArchivementDto struct {
		ArchivementId  string `json:"archivement_id" validate:"required"`
		Status         string `json:"status" validate:"required"`
		Version        int    `json:"version" validate:"required"`
		Code           string `json:"code" validate:"required"`
		Comment        string `json:"comment" validate:"required"`
		Result         string `json:"result" validate:"required"`
		IsCompileError bool   `json:"is_compile_error" validate:"required"`
	}
)

func NewArchivementController(
	RequestMapper controllers.RequestMapper,
	StudyAuthPort ports.IStudyAuthPort,
	InputPort ports.IArchivementInputPort,
) IArchivementApi {
	return &ArchivementController{
		RequestMapper: RequestMapper,
		StudyAuthPort: StudyAuthPort,
		InputPort:     InputPort,
	}
}
