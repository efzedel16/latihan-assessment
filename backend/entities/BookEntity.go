package entities

import "time"

type Book struct {
	BookId    int
	UserId    int
	Title     string
	Author    string
	Year      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
