package interactors

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/errors"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"
)

type UserInteractor struct {
	Repository ports.IUserRepository
}

func (it *UserInteractor) GetUser(ctx context.Context, userId models.UserId) (*models.User, error) {
	user, err := it.Repository.FetchByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &errors.NotFoundError{Sources: string(userId)}
	}

	return user, nil
}

// di
func NewUserInputPort(repository ports.IUserRepository) ports.IUserInputPort {
	return &UserInteractor{
		Repository: repository,
	}
}
