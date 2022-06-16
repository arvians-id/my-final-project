package model

import (
	"time"
)

type GetModuleResponse struct {
	Id       int       `json:"id,omitempty"`
	CourseId int       `json:"course_id"`
	Name     string    `json:"name"`
	IsLocked bool      `json:"is_locked"`
	Estimate int       `json:"estimate"`
	Deadline time.Time `json:"deadline"`
	Grade    *int      `json:"grade,omitempty"`
}

type CreateModuleRequest struct {
	CourseId int       `json:"course_id" binding:"required"`
	Name     string    `json:"name" binding:"required,max=50"`
	IsLocked bool      `json:"is_locked"`
	Estimate int       `json:"estimate" binding:"required"`
	Deadline time.Time `json:"deadline"`
}

type CreateModuleByCourseRequest struct {
	CodeCourse string    `json:"course_id"`
	Name       string    `json:"name" binding:"required,max=50"`
	IsLocked   bool      `json:"is_locked"`
	Estimate   int       `json:"estimate" binding:"required"`
	Deadline   time.Time `json:"deadline"`
}

type UpdateModuleRequest struct {
	CourseId int       `json:"course_id" binding:"required"`
	Name     string    `json:"name" binding:"required,max=50"`
	IsLocked bool      `json:"is_locked"`
	Estimate int       `json:"estimate" binding:"required"`
	Deadline time.Time `json:"deadline"`
}
