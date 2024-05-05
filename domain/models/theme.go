package models

type Theme struct {
	ThemeId      ThemeId
	Theme        string
	Description  string
	Archivements []*Archivement
}

// vo
type (
	ThemeId      string
	Archivements []*Archivement
)
