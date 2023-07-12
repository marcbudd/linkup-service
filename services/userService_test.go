package services_test

import (
	"testing"

	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/services"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_InvalidEmail(t *testing.T) {
	// Arrange
	req := models.UserCreateRequestDTO{
		Email:    "invalidemail",
		Username: "testuser",
		Password: "Test123!",
	}

	// Act
	user, err := services.CreateUser(req)

	// Assert
	assert.Nil(t, user)
	assert.EqualError(t, err, "email is not valid")
}

func TestCreateUser_InvalidUsername(t *testing.T) {
	// Arrange
	req := models.UserCreateRequestDTO{
		Email:    "test@example.com",
		Username: "user123!",
		Password: "Test123!",
	}

	// Act
	user, err := services.CreateUser(req)

	// Assert
	assert.Nil(t, user)
	assert.EqualError(t, err, "username is not valid")
}

func TestCreateUser_InvalidPassword(t *testing.T) {
	// Arrange
	req := models.UserCreateRequestDTO{
		Email:    "test@example.com",
		Username: "testuser",
		Password: "weak",
	}

	// Act
	user, err := services.CreateUser(req)

	// Assert
	assert.Nil(t, user)
	assert.EqualError(t, err, "password is not strong enough")
}
