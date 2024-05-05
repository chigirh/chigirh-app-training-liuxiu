package mysql

import (
	"chigirh-app-trainning-liuxiu/app/ports"
	"chigirh-app-trainning-liuxiu/domain/models"
	"context"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type RevisionGateway struct{}

type Revision struct {
	ArchivementId  string `gorm:"primaryKey;column:archivement_id"`
	Version        int    `gorm:"primaryKey;column:version"`
	Status         string `gorm:"column:status"`
	Code           string `gorm:"column:code"`
	Comment        string `gorm:"column:comment"`
	Result         string `gorm:"column:result"`
	IsCompileError bool   `gorm:"column:is_compile_error"`
}

func (it *RevisionGateway) FetchByArchivementId(
	ctx context.Context,
	archivementId models.ArchivementId,
) (*models.Revision, error) {

	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	ret := []*Revision{}
	if err := db.Where("archivement_id = ?", archivementId).Order("version desc").Limit(1).Find(&ret).Error; err != nil {
		return nil, err
	}

	if len(ret) == 0 {
		return nil, nil
	}

	e := ret[0]

	model := models.Revision{
		ArchivementId:  models.NewArchivementId(e.ArchivementId),
		Version:        models.RevisionVersion(e.Version),
		Status:         models.NewArchivementStatus(e.Status),
		Code:           models.Code(e.Code),
		Comment:        e.Comment,
		Result:         e.Result,
		IsCompileError: e.IsCompileError,
	}

	db.Close()
	return &model, nil
}

func (it *RevisionGateway) FetchByArchivementIdAndVersion(
	ctx context.Context,
	archivementId models.ArchivementId,
	version models.RevisionVersion,
) (*models.Revision, error) {

	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}

	ret := []*Revision{}
	if err := db.Where("archivement_id = ? and version = ?", archivementId, version).Find(&ret).Error; err != nil {
		return nil, err
	}

	if len(ret) == 0 {
		return nil, nil
	}

	e := ret[0]

	model := models.Revision{
		ArchivementId:  models.NewArchivementId(e.ArchivementId),
		Version:        models.RevisionVersion(e.Version),
		Status:         models.NewArchivementStatus(e.Status),
		Code:           models.Code(e.Code),
		Comment:        e.Comment,
		Result:         e.Result,
		IsCompileError: e.IsCompileError,
	}

	db.Close()
	return &model, nil
}

func (it *RevisionGateway) Create(ctx context.Context, revision models.Revision) error {
	db, err := NewDbConnection()
	if err != nil {
		return err
	}

	tx := db.Begin()

	if err := tx.Create(&Revision{
		ArchivementId:  string(revision.ArchivementId),
		Version:        int(revision.Version),
		Status:         string(revision.Status),
		Code:           string(revision.Code),
		Comment:        revision.Comment,
		Result:         revision.Result,
		IsCompileError: revision.IsCompileError,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	db.Close()
	return nil
}

func (it *RevisionGateway) UpdateByChapterIdAndVersion(ctx context.Context, revision models.Revision) error {
	db, err := NewDbConnection()
	if err != nil {
		return err
	}

	tx := db.Begin()

	err = tx.Model(&Revision{}).
		Where("archivement_id = ? and version = ?", revision.ArchivementId, revision.Version).
		Update(&Revision{
			Status:         string(revision.Status),
			Code:           string(revision.Code),
			Comment:        revision.Comment,
			Result:         revision.Result,
			IsCompileError: revision.IsCompileError,
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
func NewRevisionRepository() ports.IRevisionRepository {
	return &RevisionGateway{}
}
