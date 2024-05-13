package models

type Theme struct {
	ThemeId      ThemeId
	Theme        string
	Description  string
	Archivements []*Archivement
}

type ThemeChapters struct {
	ChapterId ChapterId
	Order     int
}

// vo
type (
	ThemeId      string
	Archivements []*Archivement
)
