package model

import "time"

type UserRegister struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Role              string    `json:"role"`
	EmailVerification time.Time `json:"email_verified_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
