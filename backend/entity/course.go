package entity

import "time"

type Course struct {
	ID          int
	Name        string
	CodeCourse  string
	Class       string
	Tools       string
	About       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
