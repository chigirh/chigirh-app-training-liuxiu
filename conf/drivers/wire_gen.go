// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package drivers

import (
	"chigirh-app-trainning-liuxiu/adapter/controllers"
	"chigirh-app-trainning-liuxiu/adapter/controllers/archivement"
	"chigirh-app-trainning-liuxiu/adapter/controllers/auth"
	"chigirh-app-trainning-liuxiu/adapter/controllers/chapter"
	"chigirh-app-trainning-liuxiu/adapter/controllers/theme"
	"chigirh-app-trainning-liuxiu/adapter/controllers/user"
	"chigirh-app-trainning-liuxiu/adapter/gateways/mysql"
	"chigirh-app-trainning-liuxiu/app/interactors"
	"context"
	"github.com/labstack/echo"
)

// Injectors from wire.go:

func InitializeDriver(ctx context.Context) (Server, error) {
	echoEcho := echo.New()
	requestMapper := controllers.NewRequestMapper()
	iAdminUserRepository := mysql.NewAdminUserRepository()
	iAdminAuthPort := interactors.NewAdminAuthPort(iAdminUserRepository)
	iAuthApi := auth.NewAuthController(requestMapper, iAdminAuthPort)
	iUserRepository := mysql.NewUserRepository()
	iUserInputPort := interactors.NewUserInputPort(iUserRepository)
	iUserApi := user.NewUserController(requestMapper, iUserInputPort)
	iStudyAuthPort := interactors.NewStudyAuthPort(iUserRepository)
	iChapterRepository := mysql.NewChapterRepository()
	iChapterInputPort := interactors.NewIChapterInputPort(iChapterRepository)
	iChapterApi := chapter.NewChapterController(requestMapper, iStudyAuthPort, iChapterInputPort)
	iThemeRepository := mysql.NewThemeRepository()
	iArchivementRepository := mysql.NewArchiveRepository()
	iThemeInputPort := interactors.NewIThemeInputPort(iThemeRepository, iArchivementRepository)
	iThemeApi := theme.NewThemeController(requestMapper, iStudyAuthPort, iThemeInputPort)
	iRevisionRepository := mysql.NewRevisionRepository()
	iArchivementInputPort := interactors.NewArchivementInputPort(iArchivementRepository, iRevisionRepository)
	iArchivementApi := archivement.NewArchivementController(requestMapper, iStudyAuthPort, iArchivementInputPort)
	server := NewDriver(echoEcho, iAuthApi, iUserApi, iChapterApi, iThemeApi, iArchivementApi)
	return server, nil
}
