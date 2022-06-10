package entity

import "time"

type Users struct {
	Id                int
	Name              string
	Username          string
	Email             string
	Password          string
	Role              int
	EmailVerification time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
