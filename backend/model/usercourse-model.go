package model

type GetUserCourseResponse struct {
	UserId   int
	CourseId int
}

type CreateUserCourseRequest struct {
	UserId   int
	CourseId int
}

type UpdateUserCourseRequest struct {
	UserId   int
	CourseId int
}
