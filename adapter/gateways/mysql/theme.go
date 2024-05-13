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

type ThemesChapterIntersection struct {
	ThemeId   string `gorm:"primaryKey;column:theme_id"`
	ChapterId string `gorm:"primaryKey;column:chapter_id"`
	Order     int    `gorm:"column:order"`
}

func (it *ThemeGateway) FetchAll(ctx context.Context) ([]*models.Theme, error) {

	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	ret := []*Theme{}
	if err := db.Find(&ret).Error; err != nil {
		return nil, err
	}

	model := []*models.Theme{}

	for _, e := range ret {
		m := models.Theme{
			ThemeId:     models.ThemeId(e.ThemeId),
			Theme:       e.Theme,
			Description: e.Description,
		}

		model = append(model, &m)
	}

	db.Close()
	return model, nil
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

func (it *ThemeGateway) Create(ctx context.Context, theme models.Theme) error {
	db, err := NewDbConnection()
	if err != nil {
		return err
	}

	tx := db.Begin()

	if err := tx.Create(&Theme{
		ThemeId:     string(theme.ThemeId),
		Theme:       string(theme.Theme),
		Description: theme.Description,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	db.Close()
	return nil
}

func (it *ThemeGateway) UpdateChapter(ctx context.Context, theme models.Theme, chapters []models.ThemeChapters) error {

	db, err := NewDbConnection()
	if err != nil {
		return err
	}

	tx := db.Begin()

	if err := tx.Where("theme_id = ?", theme.ThemeId).Delete(&ThemesChapterIntersection{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, e := range chapters {
		e := ThemesChapterIntersection{
			ThemeId:   string(theme.ThemeId),
			ChapterId: string(e.ChapterId),
			Order:     e.Order,
		}
		if err := tx.Create(&e).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	db.Close()
	return nil

}

// di
func NewThemeRepository() ports.IThemeRepository {
	return &ThemeGateway{}
}
