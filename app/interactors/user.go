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

func (it *UserInteractor) AddUser(ctx context.Context, user models.User) error {

	u, err := it.Repository.FetchByUserId(ctx, user.UserId)

	if err != nil {
		return err
	}

	if u != nil {
		return &errors.AlreadyExistsError{Sources: user.UserId}
	}

	err = it.Repository.AddUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (it *UserInteractor) GetUser(ctx context.Context, userId string) (*models.User, error) {
	u, err := it.Repository.FetchByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, &errors.NotFoundError{Sources: string(userId)}
	}

	return u, nil
}

// di
func NewUserInputPort(repository ports.IUserRepository) ports.IUserInputPort {
	return &UserInteractor{
		Repository: repository,
	}
}
