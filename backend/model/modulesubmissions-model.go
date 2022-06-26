package model

import "time"

type GetModuleSubmissionsResponse struct {
	Id          int       `json:"id,omitempty"`
	CourseId    int       `json:"course_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

type GetNextPreviousSubmissionsResponse struct {
	Id         int    `json:"id"`
	CodeCourse string `json:"code_course"`
}

type CreateModuleSubmissionsRequest struct {
	CourseId    int       `json:"course_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

type UpdateModuleSubmissionsRequest struct {
	CourseId    int       `json:"course_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}
