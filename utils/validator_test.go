package utils_test

import (
	"testing"

	"github.com/marcbudd/linkup-service/utils"
)

func TestIsValidEmail_ValidEmails(t *testing.T) {
	validEmails := []string{
		"test@example.com",
		"john.doe@example.com",
		"jane_doe@example.com",
	}

	for _, email := range validEmails {
		isValid := utils.IsValidEmail(email)

		if !isValid {
			t.Errorf("Expected email '%s' to be valid, but it was invalid", email)
		}
	}
}

func TestIsValidEmail_InvalidEmails(t *testing.T) {
	invalidEmails := []string{
		"invalid.email",
		"john.doe",
		"jane_doe@",
	}

	for _, email := range invalidEmails {
		isValid := utils.IsValidEmail(email)

		if isValid {
			t.Errorf("Expected email '%s' to be invalid, but it was valid", email)
		}
	}
}

func TestIsValidPassword_InvalidPasswords(t *testing.T) {
	invalidPasswords := []string{
		"weakpassword",
		"12345678",
		"!@#$%^&*()",
		"password123",
	}

	for _, password := range invalidPasswords {
		isValid := utils.IsValidPassword(password)

		if isValid {
			t.Errorf("Expected password '%s' to be invalid, but it was valid", password)
		}
	}
}

func TestIsValidPassword_ValidPasswords(t *testing.T) {
	validPasswords := []string{
		"SecurePassword123!",
		"AnotherStrongPassword456!",
		"thisisSecure!2232",
	}

	for _, password := range validPasswords {
		isValid := utils.IsValidPassword(password)

		if !isValid {
			t.Errorf("Expected password '%s' to be valid, but it was invalid", password)
		}
	}
}

func TestIsValidUsername_ValidUsernames(t *testing.T) {

	validUsernames := []string{
		"validUsername",
		"user_123",
		"AnotherUser42",
		"testUser_007",
	}

	for _, username := range validUsernames {
		if !utils.IsValidUsername(username) {
			t.Errorf("Erwartet: %s ist gültig, erhalten: ungültig", username)
		}
	}
}

func TestIsValidUsername_InvalidUsernames(t *testing.T) {
	invalidUsernames := []string{
		"mee",                   // Too short
		"toolongusername123456", // Too long
		"Invalid User",          // Contains spaces
		"invalid*",              // Contains unsupported special characters
		"Übername",              // Contains umlauts
	}

	for _, username := range invalidUsernames {
		if utils.IsValidUsername(username) {
			t.Errorf("Erwartet: %s ist ungültig, erhalten: gültig", username)
		}
	}
}
