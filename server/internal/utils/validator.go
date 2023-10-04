package utils

import "regexp"

func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}
