package utils_test

import (
	"testing"

	"github.com/marcbudd/linkup-service/utils"
)

func TestHashPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Errorf("Fehler beim Hashen des Passworts: %v", err)
	}

	if len(hashedPassword) == 0 {
		t.Error("Der gehashte Passwort-String ist leer.")
	}

	// Überprüfen, ob das ursprüngliche Passwort nicht identisch mit dem gehashten Passwort ist
	if password == hashedPassword {
		t.Error("Das ursprüngliche Passwort ist identisch mit dem gehashten Passwort.")
	}
}

func TestHashPasswordEmpty(t *testing.T) {
	password := ""
	_, err := utils.HashPassword(password)
	if err != nil {
		t.Error("Kein Fehler wurde erwartet, wenn das Passwort leer ist.")
	}
}

func TestHashPasswordShort(t *testing.T) {
	password := "short"
	_, err := utils.HashPassword(password)
	if err != nil {
		t.Error("Kein Fehler wurde erwartet, wenn das Passwort zu kurz ist.")
	}
}

func TestHashPasswordLong(t *testing.T) {
	password := "verylongpassword12345678901234567890123456789verylonglong"
	_, err := utils.HashPassword(password)
	if err != nil {
		t.Error("Kein Fehler wurde erwartet, wenn das Passwort zu lang ist.")
		t.Error(err)
	}
}

func TestCheckPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Errorf("Fehler beim Hashen des Passworts: %v", err)
	}

	// Check if right password returns true
	if !utils.CheckPassword(password, hashedPassword) {
		t.Error("Das überprüfte Passwort stimmt nicht mit dem gehashten Passwort überein.")
	}
}

func TestCheckPasswordInvalid(t *testing.T) {
	password := "password123"
	invalidPassword := "wrongpassword"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Errorf("Fehler beim Hashen des Passworts: %v", err)
	}

	// Check if wrong password returns false
	if utils.CheckPassword(invalidPassword, hashedPassword) {
		t.Error("Ein ungültiges Passwort wurde fälschlicherweise als gültig erkannt.")
	}
}

func TestCheckPasswordEmpty(t *testing.T) {
	password := ""
	hashedPassword, err := utils.HashPassword("password123")
	if err != nil {
		t.Errorf("Fehler beim Hashen des Passworts: %v", err)
	}

	// Check if empty password returns false
	if utils.CheckPassword(password, hashedPassword) {
		t.Error("Ein leeres Passwort wurde fälschlicherweise als gültig erkannt.")
	}
}
