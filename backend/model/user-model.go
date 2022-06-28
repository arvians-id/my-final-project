package model

import "time"

type UserRegisterResponse struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Role              int       `json:"role"`
	Phone             string    `json:"phone"`
	Gender            int       `json:"gender"`
	DisabilityType    int       `json:"type_of_disability"`
	Birthdate         string    `json:"birthdate"`
	EmailVerification time.Time `json:"email_verification"`
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
}

type GetUserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Role           int    `json:"role"`
	Gender         int    `json:"gender"`
	DisabilityType int    `json:"type_of_disability"`
}

type GetUserDetailUpdate struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Role           int       `json:"role"`
	Phone          string    `json:"phone"`
	Gender         int       `json:"gender"`
	DisabilityType int       `json:"type_of_disability"`
	Address        string    `json:"address"`
	Birthdate      string    `json:"birthdate"`
	Image          string    `jaon:"image"`
	Description    string    `json:"description"`
	UpdateAt       time.Time `json:"updated_at"`
}

type UserDetailResponse struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Username       string  `json:"username"`
	Role           int     `json:"role"`
	Phone          string  `json:"phone"`
	Gender         int     `json:"gender"`
	DisabilityType int     `json:"type_of_disability"`
	Address        *string `json:"address,omitempty"`
	Birthdate      string  `json:"birthdate"`
	Image          *string `json:"image,omitempty"`
	Description    *string `json:"description,omitempty"`
}
