package ports

import (
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

// receiving user data
type IUserInputPort interface {
	AddUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, userId string) (*models.User, error)
}

// CRUD user data to something
type IUserRepository interface {
	AddUser(ctx context.Context, user models.User) error
	FetchByUserId(ctx context.Context, userId string) (*models.User, error)
}

// Admin user
// Admin user data to
type IAdminUserRepository interface {
	FetchBy(ctx context.Context, id models.UserId, pw models.Password) (*models.AdminUser, error)
}
