package entities

import "time"

type User struct {
	UserId       int       `json:"user_id"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	Birth        time.Time `json:"birth"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
