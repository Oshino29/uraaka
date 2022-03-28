package utils

import (
	"regexp"
)

// convert string "2022-03-28 02:26:53" to string "2022-03-28 02:26"
func StripSecondsFromTime(timeString string) string {
	match := regexp.MustCompile(`( \d\d:\d\d):\d\d$`)
	replace := "$1"

	return match.ReplaceAllString(timeString, replace)
}