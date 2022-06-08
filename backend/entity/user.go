package entity

import "time"

type User struct {
	ID            int
	Name          string
	Unique_number string
	Phone         string
	Email         string
	Password      string
	Role          int
	Image         string
	Created_at    time.Time
	Updated_at    time.Time
}
