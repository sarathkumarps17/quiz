package utils

import (
	"time"
)

// ParseTimeDuration takes a string representing a time duration and
// returns the corresponding time.Duration value. If the string does not
// contain a unit of time, it is assumed to be in seconds.
func ParseTimeDuration(s string) (time.Duration, error) {
	if s == "0" {
		return time.Duration(24 * time.Hour), nil
	}
	timeout, err := time.ParseDuration(s)
	if err != nil {
		timeout_, err := time.ParseDuration(s + "s")
		if err != nil {
			return 0, err
		}
		return timeout_, nil
	}
	return timeout / time.Nanosecond, nil
}
