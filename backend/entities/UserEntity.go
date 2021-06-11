package entities

import "time"

type User struct {
	UserId       int
	Address      string
	Birth        time.Time
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
