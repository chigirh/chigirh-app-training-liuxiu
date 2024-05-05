package ports

import (
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

// receiving study user
type IUserInputPort interface {
	GetUser(ctx context.Context, userId models.UserId) (*models.User, error)
}

// repository
// study user
type IUserRepository interface {
	FetchByUserId(ctx context.Context, userId models.UserId) (*models.User, error)
	FetchBySessionKey(ctx context.Context, sessionKey models.SessionKey) (*models.User, error)
}

// Admin user
type IAdminUserRepository interface {
	FetchBy(ctx context.Context, id models.UserId, pw models.Password) (*models.AdminUser, error)
}
