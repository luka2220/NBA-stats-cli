package utils

import (
	"time"
)

func GetCurrentDate() string {
	// Get the current date
	currentTime := time.Now().UTC()

	const YYYYMMDD = "2006-01-02" // Must use this string to get the format as YYYYMMDD
	date := currentTime.Format(YYYYMMDD)

	return date
}
