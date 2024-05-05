package mysql

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ArchivementGateway struct{}

type ThemeArchivement struct {
	ArchivementId string `gorm:"primaryKey;column:archivement_id"`
	ChapterId     string `gorm:"column:chapter_id"`
	Order         int    `gorm:"column:order"`
	Status        string `gorm:"column:status"`
}

type Archivement struct {
	ArchivementId string `gorm:"primaryKey;column:archivement_id"`
	ChapterId     string `gorm:"column:chapter_id"`
	UserId        string `gorm:"column:user_id"`
	Status        string `gorm:"column:status"`
}

func (it *ArchivementGateway) FetchByUserIdAndThemeId(
	ctx context.Context,
	userId models.UserId,
	themeId models.ThemeId,
) ([]*models.Archivement, error) {

	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	ret := []*ThemeArchivement{}
	err = db.Table("themes_chapter_intersections tci").
		Select("a.archivement_id, tci.chapter_id, tci.order, a.status").
		Joins("left outer join archivements a ON a.user_id = ? and tci.chapter_id = a.chapter_id", userId).
		Where("tci.theme_id = ?", themeId).
		Scan(&ret).Error

	if err != nil {
		return nil, err
	}

	model := []*models.Archivement{}
	for _, e := range ret {
		m := models.Archivement{
			ArchivementId: models.NewArchivementId(e.ArchivementId),
			ChapterId:     models.ChapterId(e.ChapterId),
			Order:         e.Order,
			Status:        models.NewArchivementStatus(e.Status),
		}
		model = append(model, &m)
	}

	db.Close()
	return model, nil
}

func (it *ArchivementGateway) FetchByUserIdAndChapterId(
	ctx context.Context,
	userId models.UserId,
	chapterId models.ChapterId,
) (*models.Archivement, error) {

	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	ret := []*Archivement{}
	if err := db.Where("user_id = ? and chapter_id = ?", userId, chapterId).Find(&ret).Error; err != nil {
		return nil, err
	}

	if len(ret) == 0 {
		return nil, nil
	}

	e := ret[0]

	model := models.Archivement{
		ArchivementId: models.NewArchivementId(e.ArchivementId),
		ChapterId:     models.ChapterId(e.ChapterId),
		UserId:        models.UserId(e.UserId),
		Status:        models.NewArchivementStatus(e.Status),
	}

	db.Close()
	return &model, nil
}

func (it *ArchivementGateway) UpdateByArchivementId(ctx context.Context, archivement models.Archivement) error {
	db, err := NewDbConnection()
	if err != nil {
		return err
	}
	ret := []*Archivement{}
	if err := db.Where("archivement_id = ?", archivement.ArchivementId).Find(&ret).Error; err != nil {
		return err
	}

	tx := db.Begin()

	if len(ret) == 0 {
		if err := tx.Create(&Archivement{
			ArchivementId: string(archivement.ArchivementId),
			ChapterId:     string(archivement.ChapterId),
			UserId:        string(archivement.UserId),
			Status:        string(archivement.Status),
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Model(&Archivement{}).
		Where("archivement_id = ?", archivement.ArchivementId).
		Update(&Revision{
			Status: string(archivement.Status),
		}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	db.Close()
	return nil
}

// di
func NewArchiveRepository() ports.IArchivementRepository {
	return &ArchivementGateway{}
}
