package entity

import "time"

type Questions struct {
	Id          int
	CourseId    int
	UserId      int
	Title       string
	Tags        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type QuestionCourse struct {
	Id          int
	CourseId    int
	CourseName  string
	CourseClass string
	UserId      int
	UserName    string
	Title       string
	Tags        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
