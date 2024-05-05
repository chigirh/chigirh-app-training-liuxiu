//go:build wireinject
// +build wireinject

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

	"github.com/google/wire"
	"github.com/labstack/echo"
)

func InitializeDriver(ctx context.Context) (Server, error) {
	wire.Build(
		// Driver
		NewDriver,
		// echo
		echo.New,
		// commons
		controllers.NewRequestMapper,
		// controllers
		auth.NewAuthController,
		user.NewUserController,
		chapter.NewChapterController,
		theme.NewThemeController,
		archivement.NewArchivementController,
		// auth
		interactors.NewStudyAuthPort,
		interactors.NewAdminAuthPort,
		// user
		interactors.NewUserInputPort,
		mysql.NewUserRepository,
		// admin user
		mysql.NewAdminUserRepository,
		// chapter
		mysql.NewChapterRepository,
		interactors.NewIChapterInputPort,
		// theme
		mysql.NewThemeRepository,
		interactors.NewIThemeInputPort,
		// archive
		mysql.NewArchiveRepository,
		interactors.NewArchivementInputPort,
		// revision
		mysql.NewRevisionRepository,
	)
	return &Driver{}, nil
}
