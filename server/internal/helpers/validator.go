package helpers

import "regexp"

func ValidateUsername(username string) bool {
	pattern := `^[a-zA-Z0-9._-]{3,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(username)
}

func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

func ValidatePassword(password string) bool {
	pattern := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(password)
}

func ValidateDate(date string) bool {
	pattern := `^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{6}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(date)
}
