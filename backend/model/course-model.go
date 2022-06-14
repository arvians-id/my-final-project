package model

import "time"

type GetCourseResponse struct {
	Id          int       `json:"id,omitempty"`
	Name        string    `json:"name"`
	CodeCourse  string    `json:"code_course"`
	Class       string    `json:"class"`
	Tools       string    `json:"tools"`
	About       string    `json:"about"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateCourseRequest struct {
	Name        string `json:"name" binding:"required,max=50"`
	Class       string `json:"class" binding:"required,max=20"`
	Tools       string `json:"tools"`
	About       string `json:"about"`
	Description string `json:"description"`
}

type UpdateCourseRequest struct {
	Name        string `json:"name" binding:"required,max=50"`
	Class       string `json:"class" binding:"required,max=20"`
	Tools       string `json:"tools"`
	About       string `json:"about"`
	Description string `json:"description"`
}

type UpdateStatusCourseRequest struct {
	IsActive bool `json:"is_active"`
}
