package interactors

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/errors"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

type AdminAuthInteractor struct {
	Repository ports.IAdminUserRepository
}

func (it *AdminAuthInteractor) AuthAdminUser(ctx context.Context, id models.UserId, pw models.Password) (bool, error) {

	user, err := it.Repository.FetchBy(ctx, id, pw)

	if err != nil {
		return false, err
	}

	if user == nil {
		return false, &errors.AuthenticationError{Sources: string(id)}
	}

	return true, nil
}

// di
func NewAdminAuthPort(repository ports.IAdminUserRepository) ports.IAdminAuthPort {
	return &AdminAuthInteractor{
		Repository: repository,
	}
}
