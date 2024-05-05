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
}

type UserController struct {
	RequestMapper controllers.RequestMapper
	InputPort     ports.IUserInputPort
}

func (it *UserController) Get(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {
		userId := c.Param("userId")

		u, err := it.InputPort.GetUser(ctx, models.UserId(userId))

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		res := new(GetResponse)
		res.User = UserDto{
			UserId:     string(u.UserId),
			SessionKey: string(u.SessionKey),
			ThemeId:    string(u.ThemeId),
		}

		return c.JSON(http.StatusOK, res)
	}
}

// dto -->
type (
	GetResponse struct {
		User UserDto `json:"user" validate:"required"`
	}

	UserDto struct {
		UserId     string `json:"user_id" validate:"required,max=64"`
		SessionKey string `json:"session_key" validate:"required,max=64"`
		ThemeId    string `json:"theme_id" validate:"required,max=64"`
	}
)

func NewUserController(
	RequestMapper controllers.RequestMapper,
	InputPort ports.IUserInputPort,
) IUserApi {
	return &UserController{
		RequestMapper: RequestMapper,
		InputPort:     InputPort,
	}
}
