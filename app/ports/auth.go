package ports

import (
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

// receiving admin data
type IAdminAuthPort interface {
	AuthAdminUser(ctx context.Context, id models.UserId, pw models.Password) (bool, error)
}
