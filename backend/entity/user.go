package entity

import "time"

type User struct {
	ID                int
	Name              string
	UserName          string
	Email             string
	Password          string
	Role              string
	EmailVerification time.Time
	Created_at        time.Time
	Updated_at        time.Time
}
