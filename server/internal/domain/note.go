package domain

import "time"

type Note struct {
	Id        string
	Title     string
	AuthorID  string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
