package services_test

import (
	"testing"

	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/services"
	"github.com/stretchr/testify/assert"
)

// func TestCreateUser_Success(t *testing.T) {
// 	// Arrange
// 	req := models.UserCreateRequestDTO{
// 		Email:    "test@example.com",
// 		Username: "testuser",
// 		Password: "Test123!",
// 	}

// 	// TODO: Create mock database

// 	// Act
// 	user, err := services.CreateUser(req)

// 	// Assert
// 	assert.Nil(t, err)
// 	assert.NotNil(t, user)
// 	assert.Equal(t, req.Email, user.Email)
// 	assert.Equal(t, req.Username, user.Username)
// }

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

// func TestCreateUser_DuplicateEmail(t *testing.T) {
// 	// Arrange
// 	req := models.UserCreateRequestDTO{
// 		Email:    "existing@example.com",
// 		Username: "testuser",
// 		Password: "Test123!",
// 	}
// 	// TODO: Simulate existing email in the database

// 	// Act
// 	user, err := services.CreateUser(req)

// 	// Assert
// 	assert.Nil(t, user)
// 	assert.EqualError(t, err, "email is already taken")
// }

// func TestCreateUser_DuplicateUsername(t *testing.T) {
// 	// Arrange
// 	req := models.UserCreateRequestDTO{
// 		Email:    "test@example.com",
// 		Username: "existinguser",
// 		Password: "Test123!",
// 	}
// 	// TODO: Simulate existing username in the database

// 	// Act
// 	user, err := services.CreateUser(req)

// 	// Assert
// 	assert.Nil(t, user)
// 	assert.EqualError(t, err, "username is already taken")
// }
