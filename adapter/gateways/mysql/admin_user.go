package mysql

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AdminUserGateway struct{}

type AdminUser struct {
	UserId string `gorm:"primaryKey;column:user_id"`
}

func (it *AdminUserGateway) FetchBy(ctx context.Context, id models.UserId, pw models.Password) (*models.AdminUser, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	ret := []*AdminUser{}
	if err := db.Where("user_id = ? and password = ? ", id, pw).Find(&ret).Error; err != nil {
		return nil, err
	}

	if len(ret) == 0 {
		return nil, nil
	}

	e := ret[0]

	model := models.AdminUser{
		Id: models.UserId(e.UserId),
	}

	db.Close()
	return &model, nil
}

// di
func NewAdminUserRepository() ports.IAdminUserRepository {
	return &AdminUserGateway{}
}
