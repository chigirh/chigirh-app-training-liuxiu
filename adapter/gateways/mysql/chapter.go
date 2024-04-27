package mysql

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ChapterGateway struct{}

type Chapter struct {
	ChapterId       string `gorm:"primaryKey;column:chapter_id"`
	MainExecuteCode string `gorm:"column:main_execute_code"`
	InitCode        string `gorm:"column:init_code"`
	Expected        string `gorm:"column:expected"`
	AnswerCode      string `gorm:"column:answer_code"`
	Level           int    `gorm:"column:level"`
}

func (it *ChapterGateway) Create(ctx context.Context, chapter models.Chapter) error {
	db, err := NewDbConnection()
	if err != nil {
		return err
	}

	tx := db.Begin()

	// users
	if err := tx.Create(&Chapter{
		ChapterId:       string(chapter.Id),
		MainExecuteCode: string(chapter.MainExecute),
		InitCode:        string(chapter.Init),
		Expected:        chapter.Expected,
		AnswerCode:      string(chapter.Answer),
		Level:           int(chapter.Level),
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
		Id:          models.ChapterId(e.ChapterId),
		MainExecute: models.Code(e.MainExecuteCode),
		Init:        models.Code(e.InitCode),
		Expected:    e.Expected,
		Answer:      models.Code(e.AnswerCode),
		Level:       models.Level(e.Level),
	}

	db.Close()
	return &model, nil
}

// di
func NewChapterRepository() ports.IChapterRepository {
	return &ChapterGateway{}
}
