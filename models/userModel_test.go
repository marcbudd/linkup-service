package models_test

import (
	"testing"
	"time"

	"github.com/marcbudd/linkup-service/models"
)

func TestUpdateUser(t *testing.T) {
	user := &models.User{}
	req := models.UserUpdateRequestDTO{
		Username:  "newusername",
		BirthDate: time.Now(),
		Name:      "New Name",
		Bio:       "New Bio",
		Image:     nil,
	}

	user.UpdateUser(req)

	if user.Username != req.Username {
		t.Errorf("Expected username %s, but got %s", req.Username, user.Username)
	}

	if user.BirthDate != req.BirthDate {
		t.Errorf("Expected birth date %s, but got %s", req.BirthDate, user.BirthDate)
	}

	if user.Name != req.Name {
		t.Errorf("Expected name %s, but got %s", req.Name, user.Name)
	}

	if user.Bio != req.Bio {
		t.Errorf("Expected bio %s, but got %s", req.Bio, user.Bio)
	}

	if user.Image != req.Image {
		t.Errorf("Expected image %v, but got %v", req.Image, user.Image)
	}
}

func TestConvertUserToResponseDTO(t *testing.T) {
	user := &models.User{
		ID:        1,
		Username:  "username",
		BirthDate: time.Now(),
		Name:      "John Doe",
		Bio:       "Lorem ipsum",
		Image:     nil,
	}

	response := user.ConvertUserToResponseDTO()

	if response.ID != user.ID {
		t.Errorf("Expected ID %d, but got %d", user.ID, response.ID)
	}

	if response.Username != user.Username {
		t.Errorf("Expected username %s, but got %s", user.Username, response.Username)
	}

	if response.BirthDate != user.BirthDate {
		t.Errorf("Expected birth date %s, but got %s", user.BirthDate, response.BirthDate)
	}

	if response.Name != user.Name {
		t.Errorf("Expected name %s, but got %s", user.Name, response.Name)
	}

	if response.Bio != user.Bio {
		t.Errorf("Expected bio %s, but got %s", user.Bio, response.Bio)
	}

	if response.Image != user.Image {
		t.Errorf("Expected image %v, but got %v", user.Image, response.Image)
	}
}
