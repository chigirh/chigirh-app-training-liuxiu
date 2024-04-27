package auth

import (
	"context"
	"net/http"

	"chigirh-app-trainning-liuxiu/adapter/controllers"
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/conf/config"
	"chigirh-app-trainning-liuxiu/domain/models"

	"github.com/labstack/echo"
)

type IAuthApi interface {
	AdminAuth(ctx context.Context) func(c echo.Context) error
}

type AuthController struct {
	requestMapper controllers.RequestMapper
	adminAuthPort ports.IAdminAuthPort
}

func (it *AuthController) AdminAuth(ctx context.Context) func(c echo.Context) error {

	return func(c echo.Context) error {

		req := new(AdminAuthRequest)
		if err := it.requestMapper.Parse(c, req); err != nil {
			return err
		}

		isAuthorized, err := it.adminAuthPort.AuthAdminUser(ctx, models.UserId(req.Id), models.Password(req.Password))

		if err != nil {
			return controllers.ErrorHandle(c, err)
		}

		if !isAuthorized {
			return c.JSON(http.StatusUnauthorized, controllers.DefaultResponse)
		}

		res := new(AdminAuthResponse)
		res.Key = config.Server.MasterKey

		return c.JSON(http.StatusOK, res)
	}
}

// dto -->
type (
	AdminAuthRequest struct {
		Id       string `json:"user_id" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	AdminAuthResponse struct {
		Key string `json:"master_key"`
	}
)

func NewAuthController(
	requestMapper controllers.RequestMapper,
	adminAuthPort ports.IAdminAuthPort,
) IAuthApi {
	return &AuthController{
		requestMapper: requestMapper,
		adminAuthPort: adminAuthPort,
	}
}
