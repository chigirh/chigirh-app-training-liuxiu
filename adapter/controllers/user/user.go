package user

import (
	"context"
	"net/http"

	"chigirh-app-trainning-liuxiu/adapter/controllers"
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/models"

	"github.com/labstack/echo"
)

type IUserApi interface {
	Get(ctx context.Context) func(c echo.Context) error
	Post(ctx context.Context) func(c echo.Context) error
}

type UserController struct {
	requestMapper controllers.RequestMapper
	inputPort     ports.IUserInputPort
}

func (it *UserController) Get(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {
		userId := c.Param("userId")

		// If session token is set, have admin.
		// _, err := it.requestMapper.GetSessionToken(c)
		// if err != nil {
		// 	return err
		// }

		u, err := it.inputPort.GetUser(ctx, userId)

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		res := new(GetResponse)
		res.User = UserDto{
			UserId: string(u.UserId),
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (it *UserController) Post(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {

		req := new(PostRequest)

		if err := it.requestMapper.Parse(c, req); err != nil {
			return err
		}

		user := models.User{
			UserId: req.User.UserId,
		}
		if err := it.inputPort.AddUser(ctx, user); err != nil {
			return controllers.ErrorHandle(c, err)
		}

		res := controllers.DefaultResponse
		return c.JSON(http.StatusOK, res)
	}
}

// dto -->
type (
	GetResponse struct {
		User UserDto `json:"user" validate:"required"`
	}

	PostRequest struct {
		User UserDto `json:"user" validate:"required"`
	}

	UserDto struct {
		UserId string `json:"user_id" validate:"required,max=64"`
	}
)

func NewUserController(
	requestMapper controllers.RequestMapper,
	inputPost ports.IUserInputPort,
) IUserApi {
	return &UserController{
		requestMapper: requestMapper,
		inputPort:     inputPost,
	}
}
