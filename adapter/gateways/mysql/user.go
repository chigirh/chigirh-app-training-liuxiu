package mysql

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserGateway struct{}

type User struct {
	UserId     string `gorm:"primaryKey;column:user_id"`
	SessionKey string `gorm:"column:session_key"`
	ThemeId    string `gorm:"column:theme_id"`
}

func (it *UserGateway) FetchByUserId(ctx context.Context, userId models.UserId) (*models.User, error) {

	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	ret := []*User{}
	if err := db.Where("user_id = ?", userId).Find(&ret).Error; err != nil {
		return nil, err
	}

	if len(ret) == 0 {
		return nil, nil
	}

	e := ret[0]

	model := models.User{
		UserId:     models.UserId(e.UserId),
		SessionKey: models.SessionKey(e.SessionKey),
		ThemeId:    models.ThemeId(e.ThemeId),
	}

	db.Close()
	return &model, nil
}

func (it *UserGateway) FetchBySessionKey(ctx context.Context, sessionKey models.SessionKey) (*models.User, error) {
	// user
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	userResult := []*User{}
	if err := db.Where("session_key = ?", sessionKey).Find(&userResult).Error; err != nil {
		return nil, err
	}

	if len(userResult) == 0 {
		return nil, nil
	}

	e := userResult[0]

	model := models.User{
		UserId:     models.UserId(e.UserId),
		SessionKey: models.SessionKey(e.SessionKey),
		ThemeId:    models.ThemeId(e.ThemeId),
	}

	db.Close()
	return &model, nil
}

// di
func NewUserRepository() ports.IUserRepository {
	return &UserGateway{}
}
