package interactors

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/errors"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

type StudyAuthInteractor struct {
	Repository ports.IUserRepository
}

func (it *StudyAuthInteractor) GetAuthorizedUser(ctx context.Context, sessionKey models.SessionKey) (*models.User, error) {
	user, err := it.Repository.FetchBySessionKey(ctx, sessionKey)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &errors.AuthorizationError{Sources: string(sessionKey)}
	}

	return user, nil
}

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
func NewStudyAuthPort(repository ports.IUserRepository) ports.IStudyAuthPort {
	return &StudyAuthInteractor{
		Repository: repository,
	}
}

func NewAdminAuthPort(repository ports.IAdminUserRepository) ports.IAdminAuthPort {
	return &AdminAuthInteractor{
		Repository: repository,
	}
}
