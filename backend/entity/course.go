package entity

import "time"

type Course struct {
	ID          int
	Name        string
	Code_course string
	Class       string
	Tools       string
	About       string
	Description string
	Created_at  time.Time
	Updated_at  time.Time
}
