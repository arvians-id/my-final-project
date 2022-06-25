package model

import "time"

type GetAnswerResponse struct {
	Id          int 			`json:"id,omitempty"`
	QuestionId  int 			`json:"question_id"`
	UserId 		  int 			`json:"user_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateAnswerRequest struct {
	QuestionId  int 			`json:"question_id"`
	UserId 		  int 			`json:"user_id"`
	Description string    `json:"description"`
}

type UpdateAnswerRequest struct {
	QuestionId  int 			`json:"question_id"`
	UserId 		  int 			`json:"user_id"`
	Description string    `json:"description"`
}
