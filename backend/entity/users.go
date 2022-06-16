package entity

import "time"

type Users struct {
	Id                int
	Name              string
	Username          string
	Email             string
	Password          string
	Role              int
	Phone             string
	Gender            int
	DisabilityType    int
	Address           string
	Birthdate         string
	Image             string
	Description       string
	EmailVerification time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
