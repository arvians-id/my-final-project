package entity

import (
	"time"
)

type Modules struct {
	Id       int
	CourseId int
	Name     string
	IsLocked bool
	Estimate int
	Deadline time.Time
	Grade    *int
}
