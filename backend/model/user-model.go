package model

import "time"

type UserRegister struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Unique_number string    `json:"unique_number"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Role          int       `json:"role"`
	Image         string    `json:"image"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}
