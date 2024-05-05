package models

type User struct {
	UserId     UserId
	SessionKey SessionKey
	ThemeId    ThemeId
}

type AdminUser struct {
	Id UserId
}

// vo
type (
	UserId   string
	Password string
)
