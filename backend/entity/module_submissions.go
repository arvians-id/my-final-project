package entity

import "time"

type ModuleSubmissions struct {
	Id          int
	CourseId    int
	Name        string
	Description string
	Deadline    time.Time
}

type NextPreviousModuleSubmissions struct {
	Id         int
	CodeCourse string
}
