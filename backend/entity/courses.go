package entity

import "time"

type Courses struct {
	Id          int
	Name        string
	CodeCourse  string
	Class       string
	Tools       string
	About       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
