package utils

import "time"

func TimeNow() time.Time {
	format := "2006-01-02 15:04:05"
	timeNow, _ := time.Parse(format, time.Now().Format(format))
	return timeNow
}

func ParseTime(times time.Time) time.Time {
	timeNow, _ := time.Parse(time.RFC3339, times.Format(time.RFC3339))
	return timeNow
}
