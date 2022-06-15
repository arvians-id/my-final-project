package model

import "time"

type UserRegister struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Role              int       `json:"role"`
	Gender            int       `json:"gender"`
	DisabilityType    int       `json:"type_of_disability"`
	Birthdate         string    `json:"birthdate"`
	EmailVerification time.Time `json:"email_verification"`
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
}
