package model

type GetUserCourseResponse struct {
	UserId   int `json:"user_id"`
	CourseId int `json:"course_id"`
}

type CreateUserCourseRequest struct {
	UserId   int `json:"user_id"`
	CourseId int `json:"course_id"`
}

type GetStudentCourseResponse struct {
	IdCourse    int    `json:"id_course"`
	CourseName  string `json:"course_name"`
	CourseCode  string `json:"course_code"`
	CourseClass string `json:"course_class"`
}

type GetUserTeacherCourseResponse struct {
	IdUser       int    `json:"id_user"`
	UserName     string `json:"user_name"`
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
}

type GetStudentSubmissionsResponse struct {
	IdModuleSubmission   int     `json:"id_module_submission"`
	NameCourse           string  `json:"name_course"`
	CodeCourse           string  `json:"code_course"`
	NameModuleSubmission string  `json:"name_module_submission"`
	Grade                *int    `json:"grade,omitempty"`
	File                 *string `json:"file,omitempty"`
}

type GetTeacherSubmissionsResponse struct {
	IdUserSubmission     *int    `json:"id_user_submission,omitempty"`
	UserName             string  `json:"user_name"`
	ModuleSubmissionName string  `json:"module_submission_name"`
	Grade                *int    `json:"grade,omitempty"`
	File                 *string `json:"file,omitempty"`
}
