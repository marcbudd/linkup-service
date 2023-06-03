package utils

import (
	"regexp"
	"strings"
	"unicode"
)

func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

func IsValidPassword(password string) bool {
	// length at least 8 chars
	if len(password) < 8 {
		return false
	}

	// Check for uppercase and lower case
	hasUppercase := false
	hasLowercase := false
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUppercase = true
		}
		if char >= 'a' && char <= 'z' {
			hasLowercase = true
		}
	}
	if !hasUppercase || !hasLowercase {
		return false
	}

	// Check for number
	hasNumber := false
	for _, char := range password {
		if unicode.IsNumber(char) {
			hasNumber = true
			break
		}
	}
	if !hasNumber {
		return false
	}

	// Check for special character
	hasSpecialChar := strings.ContainsAny(password, "!@#$%^&*()_+-=[]{};:'\"\\|,.<>/?")
	return hasSpecialChar

}

func IsValidUsername(username string) bool {
	// Check: is username is between 4 and 16 characters long
	if len(username) < 4 || len(username) > 16 {
		return false
	}

	// Check: username contains only alphanumeric characters and underscores
	validUsername := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return validUsername.MatchString(username)
}
