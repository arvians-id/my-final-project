package entity

type UserCourse struct {
	UserId   int
	CourseId int
}

type StudentCourse struct {
	IdCourse    int
	CourseName  string
	CourseCode  string
	CourseClass string
}

type UserTeacherCourse struct {
	IdUser       int
	UserName     string
	UserUsername string
	UserEmail    string
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
