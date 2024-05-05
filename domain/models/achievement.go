package models

import (
	"github.com/google/uuid"
)

type Archivement struct {
	ArchivementId ArchivementId
	ChapterId     ChapterId
	UserId        UserId
	Order         int
	Status        ArchivementStatus
	Revision      Revision
}

// vo
type (
	ArchivementId     string
	ArchivementStatus string
)

// constractor
func NewArchivementId(v string) ArchivementId {
	if v == "" {
		uuid, _ := uuid.NewRandom()
		return ArchivementId(uuid.String())
	}
	return ArchivementId(v)
}

func NewArchivementStatus(v string) ArchivementStatus {
	if v == "" {
		return ArchivementStatus("0")
	}
	return ArchivementStatus(v)
}
