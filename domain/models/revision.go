package models

type Revision struct {
	ArchivementId  ArchivementId
	Version        RevisionVersion
	Status         ArchivementStatus
	Code           Code
	Comment        string
	Result         string
	IsCompileError bool
}

// vo
type (
	RevisionVersion int
)

func NewRevisionVersion(v int) RevisionVersion {
	if v == 0 {
		return RevisionVersion(1)
	}
	return RevisionVersion(v)
}
