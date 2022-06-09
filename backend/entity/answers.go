package entity

import "time"

type Answers struct {
	Id          int
	QuestionId  int
	UserId      int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
