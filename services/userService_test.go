package services_test

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/services"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreateUser_Success(t *testing.T) {
	// Arrange
	email := "john.doe@example.com"
	password := "strongPassword123!"
	username := "johndoe"
	birthDate := time.Now()
	name := "John Doe"
	bio := "Hello, I'm John Doe"

	req := models.UserCreateRequestDTO{
		Email:     email,
		Username:  username,
		Password:  password,
		BirthDate: birthDate,
		Name:      name,
		Bio:       bio,
		Image:     nil,
	}

	// TODO: Create mock database
	// initalizers.DB, _ = NewMock()

	// Act
	user, err := services.CreateUser(req)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, birthDate, user.BirthDate)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, bio, user.Bio)
	assert.Nil(t, user.Image)

}

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
