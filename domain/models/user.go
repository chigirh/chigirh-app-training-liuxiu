package models

type User struct {
	UserId string
}

type AdminUser struct {
	Id UserId
}

// vo
type (
	UserId   string
	Password string
)
