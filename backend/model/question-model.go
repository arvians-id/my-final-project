package model

import "time"

type GetQuestionResponse struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	CourseId    int       `json:"course_id"`
	Title       string    `json:"title"`
	Tags        string    `json:"tags"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetQuestionRelationResponse struct {
	Id          int       `json:"id"`
	CourseId    int       `json:"course_id"`
	CourseName  string    `json:"course_name"`
	CourseClass string    `json:"course_class"`
	UserId      int       `json:"user_id"`
	UserName    string    `json:"user_name"`
	Title       string    `json:"title"`
	Tags        string    `json:"tags"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateQuestionRequest struct {
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	CourseId    int    `json:"course_id"`
	Tags        string `json:"tags"`
	Description string `json:"description"`
}

type UpdateQuestionRequest struct {
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	CourseId    int    `json:"course_id"`
	Tags        string `json:"tags"`
	Description string `json:"description"`
}
