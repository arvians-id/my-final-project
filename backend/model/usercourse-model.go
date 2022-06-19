package model

type GetUserCourseResponse struct { 
	UserId   int `json:"user_id"`
	CourseId int `json:"course_id"`
}

type CreateUserCourseRequest struct {
	UserId   int `json:"user_id"`
	CourseId int `json:"course_id"`
}

type UpdateUserCourseRequest struct {
	UserId   int `json:"user_id"`
	CourseId int `json:"course_id"`
}
