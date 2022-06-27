package entity

type UserSubmissions struct {
	Id                 int
	UserId             int
	ModuleSubmissionId int
	File               *string
	Grade              *int
}
