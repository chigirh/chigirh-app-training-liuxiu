package models

type Chapter struct {
	Id          ChapterId
	MainExecute Code
	Init        Code
	Expected    string
	Answer      Code
	Level       Level
}

type (
	ChapterId string
	Code      string
	Level     int
)
