package model

type GetModuleArticlesResponse struct {
	Id       int    `json:"id,omitempty"`
	CourseId int    `json:"course_id"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	Estimate int    `json:"estimate"`
}

type GetNextPreviousArticlesResponse struct {
	Id         int    `json:"id"`
	CodeCourse string `json:"code_course"`
}

type CreateModuleArticlesRequest struct {
	CourseId int    `json:"course_id"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	Estimate int    `json:"estimate"`
}
type UpdateModuleArticlesRequest struct {
	Name     string `json:"name"`
	Content  string `json:"content"`
	Estimate int    `json:"estimate"`
}
