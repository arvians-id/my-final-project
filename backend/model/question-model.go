package model

import "time"

type GetQuestionResponse struct {
	Id						int       `json:"id"`
	UserId     		int   		`json:"user_id"`
	ModuleId     	int   		`json:"module_id"`	
	Title         string    `json:"title"`
	Tags          string    `json:"tags"`
	Description   string    `json:"description"`
	CreatedAt    time.Time 	`json:"created_at"`
	UpdatedAt    time.Time 	`json:"updated_at"`
}

type CreateQuestionRequest struct {
	UserId     		int   	`json:"user_id"`
	Title  				string	`json:"title"`
	ModuleId     	int   	`json:"module_id"`	
	Tags        	string  `json:"tags"`
	Description 	string  `json:"description"`
}


type UpdateQuestionRequest struct {
	UserId     		int   	`json:"user_id"`
	Title  				string	`json:"title"`
	ModuleId     	int   	`json:"module_id"`	
	Tags        	string  `json:"tags"`
	Description 	string  `json:"description"`
}