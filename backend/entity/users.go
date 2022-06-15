package entity

import "time"

type Users struct {
	Id                int
	Name              string
	Username          string
	Email             string
	Password          string
	Role              int
	Gender            int
	DisabilityType    int
	Birthdate         string
	EmailVerification time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
