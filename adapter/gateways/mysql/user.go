package mysql

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserGateway struct{}

type User struct {
	UserId string `json:"user_id"`
}

func (it *UserGateway) AddUser(ctx context.Context, user models.User) error {

	db, err := NewDbConnection()
	if err != nil {
		return err
	}

	tx := db.Begin()

	// users
	if err := tx.Create(&User{
		UserId: user.UserId,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	db.Close()
	return nil
}

func (it *UserGateway) FetchByUserId(ctx context.Context, userId string) (*models.User, error) {
	// user
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	userResult := []*User{}
	if err := db.Where("user_id = ?", userId).Find(&userResult).Error; err != nil {
		return nil, err
	}

	if len(userResult) == 0 {
		return nil, nil
	}

	entity := userResult[0]

	model := models.User{
		UserId: entity.UserId,
	}

	db.Close()
	return &model, nil

}

// di
func NewUserRepository() ports.IUserRepository {
	return &UserGateway{}
}
