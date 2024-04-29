package models

type Chapter struct {
	Id           ChapterId
	Main         Code
	Example      Code
	Expected     string
	BestPractice Code
	Level        Level
	Exercise     string
}

type (
	ChapterId string
	Code      string
	Level     int
)
