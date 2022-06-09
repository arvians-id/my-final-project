package entity

import "time"

type Questions struct {
	Id          int
	ModuleId    int
	UserId      int
	Title       string
	Tags        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
