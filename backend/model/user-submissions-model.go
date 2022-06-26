package model

type GetUserSubmissionsResponse struct {
	Id                 int    `json:"id,omitempty"`
	UserId             int    `json:"user_id"`
	ModuleSubmissionId int    `json:"module_submission_id"`
	File               string `json:"file"`
	Grade              *int   `json:"grade,omitempty"`
}

type CreateUserSubmissionsRequest struct {
	UserId             int
	ModuleSubmissionId int
	File               string
}

type UpdateUserGradeRequest struct {
	Id    int
	Grade int `json:"grade"`
}
