package drivers

import (
	"chigirh-app-trainning-liuxiu/adapter/controllers/auth"
	"chigirh-app-trainning-liuxiu/adapter/controllers/chapter"
	"chigirh-app-trainning-liuxiu/adapter/controllers/user"
	"chigirh-app-trainning-liuxiu/conf/config"
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server interface {
	Start(ctx context.Context)
}

type Driver struct {
	echo       *echo.Echo
	authApi    auth.IAuthApi
	userApi    user.IUserApi
	chapterApi chapter.IChapterApi
}

func NewDriver(
	echo *echo.Echo,
	authApi auth.IAuthApi,
	userApi user.IUserApi,
	chapterApi chapter.IChapterApi,
) Server {
	return &Driver{
		echo:       echo,
		authApi:    authApi,
		userApi:    userApi,
		chapterApi: chapterApi,
	}
}

func (driver *Driver) Start(ctx context.Context) {
	log.Println("api start.")
	// cors
	driver.echo.Use(middleware.CORS())
	// custom validator
	// driver.echo.Validator = controllers.NewValidator()

	// auth
	driver.echo.POST("/admin/authentication", driver.authApi.AdminAuth(ctx))
	// users
	driver.echo.GET("/users/:userId", driver.userApi.Get(ctx))
	driver.echo.POST("/users", driver.userApi.Post(ctx))
	// chapter
	driver.echo.GET("/chapter/:chapterId", driver.chapterApi.Get(ctx))
	driver.echo.POST("/chapter", driver.chapterApi.Post(ctx))

	driver.echo.Logger.Fatal(driver.echo.Start(fmt.Sprintf(":%d", config.Server.ServerPort)))
}
