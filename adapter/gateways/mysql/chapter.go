package mysql

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ChapterGateway struct{}

type Chapter struct {
	ChapterId        string `gorm:"primaryKey;column:chapter_id"`
	MainCode         string `gorm:"column:main_code"`
	ExampleCode      string `gorm:"column:example_code"`
	Expected         string `gorm:"column:expected"`
	BestPracticeCode string `gorm:"column:best_practice_code"`
	Level            int    `gorm:"column:level"`
	Exercise         string `gorm:"column:exercise"`
}

func (it *ChapterGateway) Create(ctx context.Context, chapter models.Chapter) error {
	db, err := NewDbConnection()
	if err != nil {
		return err
	}

	tx := db.Begin()

	// users
	if err := tx.Create(&Chapter{
		ChapterId:        string(chapter.Id),
		MainCode:         string(chapter.Main),
		ExampleCode:      string(chapter.Example),
		Expected:         chapter.Expected,
		BestPracticeCode: string(chapter.BestPractice),
		Level:            int(chapter.Level),
		Exercise:         chapter.Exercise,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	db.Close()
	return nil
}

func (it *ChapterGateway) FetchBy(ctx context.Context, id models.ChapterId) (*models.Chapter, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	ret := []*Chapter{}
	if err := db.Where("chapter_id = ?", id).Find(&ret).Error; err != nil {
		return nil, err
	}

	if len(ret) == 0 {
		return nil, nil
	}

	e := ret[0]

	model := models.Chapter{
		Id:           models.ChapterId(e.ChapterId),
		Main:         models.Code(e.MainCode),
		Example:      models.Code(e.ExampleCode),
		Expected:     e.Expected,
		BestPractice: models.Code(e.BestPracticeCode),
		Level:        models.Level(e.Level),
		Exercise:     e.Exercise,
	}

	db.Close()
	return &model, nil
}

func (it *ChapterGateway) FetchAll(ctx context.Context) ([]*models.Chapter, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	ret := []*Chapter{}
	if err := db.Find(&ret).Error; err != nil {
		return nil, err
	}

	if len(ret) == 0 {
		return []*models.Chapter{}, nil
	}

	model := []*models.Chapter{}

	for _, e := range ret {
		m := models.Chapter{
			Id:           models.ChapterId(e.ChapterId),
			Main:         models.Code(e.MainCode),
			Example:      models.Code(e.ExampleCode),
			Expected:     e.Expected,
			BestPractice: models.Code(e.BestPracticeCode),
			Level:        models.Level(e.Level),
			Exercise:     e.Exercise,
		}

		model = append(model, &m)
	}

	db.Close()
	return model, nil
}

func (it *ChapterGateway) FetchByThemeId(ctx context.Context, id models.ThemeId) ([]*models.Chapter, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	ret := []*Chapter{}
	err = db.Table("themes_chapter_intersections tci").
		Select("c.*").
		Joins("inner join chapters c ON c.chapter_id =  tci.chapter_id").
		Where("tci.theme_id = ?", id).
		Scan(&ret).Error

	if err != nil {
		return nil, err
	}

	if len(ret) == 0 {
		return []*models.Chapter{}, nil
	}

	model := []*models.Chapter{}

	for _, e := range ret {
		m := models.Chapter{
			Id:           models.ChapterId(e.ChapterId),
			Main:         models.Code(e.MainCode),
			Example:      models.Code(e.ExampleCode),
			Expected:     e.Expected,
			BestPractice: models.Code(e.BestPracticeCode),
			Level:        models.Level(e.Level),
			Exercise:     e.Exercise,
		}

		model = append(model, &m)
	}

	db.Close()
	return model, nil
}

// di
func NewChapterRepository() ports.IChapterRepository {
	return &ChapterGateway{}
}
