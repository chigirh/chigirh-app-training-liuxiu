package ports

import (
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

// receiving admin
type IStudyAuthPort interface {
	GetAuthorizedUser(ctx context.Context, sessionKey models.SessionKey) (*models.User, error)
}

type IAdminAuthPort interface {
	AuthAdminUser(ctx context.Context, id models.UserId, pw models.Password) (bool, error)
}
