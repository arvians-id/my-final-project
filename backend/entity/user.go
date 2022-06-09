package entity

import "time"

type User struct {
	Id                int
	Name              string
	UserName          string
	Email             string
	Password          string
	Role              string
	EmailVerification time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
