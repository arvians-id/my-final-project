package entity

import (
	"time"
)

const (
	TimeFormat = "2006-01-02"
)

type User_details struct {
	ID                 int
	User_id            int
	Phone              string
	Gender             int
	Type_of_disability int
	Address            string
	Birthdate          Birthdate
	Image              string
	Description        string
}

type Birthdate interface {
	getDOB(year, month, day int) time.Time
}

func getDOB(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}

/*
	Reference:
	https://golangbyexample.com/dob-golang/

	format of date:
	years-months-days
*/
