package week

import "time"

// Get the current week number.
func Now() int {
	t := time.Now()
	_, week := t.ISOWeek()
	return week
}
