package entity

type UserCourse struct {
	UserId   int
	CourseId int
}

type StudentSubmissions struct {
	IdModuleSubmission   int
	CourseName           string
	ModuleSubmissionName string
	Grade                *int
	File                 *string
}

type TeacherSubmissions struct {
	IdUserSubmission     *int
	UserName             string
	ModuleSubmissionName string
	Grade                *int
	File                 *string
}
