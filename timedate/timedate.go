package timedate

import "time"

// SameDay return a bool if two dates are
// the same year and same year day.
func SameDay(a, b time.Time) bool {
	return a.Year() == b.Year() &&
		a.YearDay() == b.YearDay()
}
