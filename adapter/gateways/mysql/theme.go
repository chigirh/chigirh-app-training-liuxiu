package mysql

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ThemeGateway struct{}

type Theme struct {
	ThemeId     string `gorm:"column:theme_id"`
	Theme       string `gorm:"column:theme"`
	Description string `gorm:"column:description"`
}

func (it *ThemeGateway) FetchByThemeId(ctx context.Context, themeId models.ThemeId) (*models.Theme, error) {

	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	ret := []*Theme{}
	if err := db.Where("theme_id = ?", themeId).Find(&ret).Error; err != nil {
		return nil, err
	}

	if len(ret) == 0 {
		return nil, nil
	}

	e := ret[0]

	model := models.Theme{
		ThemeId:     models.ThemeId(e.ThemeId),
		Theme:       e.Theme,
		Description: e.Description,
	}

	db.Close()
	return &model, nil
}

// di
func NewThemeRepository() ports.IThemeRepository {
	return &ThemeGateway{}
}
