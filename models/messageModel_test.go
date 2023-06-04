package models_test

import (
	"testing"
	"time"

	"github.com/marcbudd/linkup-service/models"
)

func TestConvertMessageToResponseDTO(t *testing.T) {
	currentUserID := uint(1)
	chatPartnerID := uint(2)
	message := &models.Message{
		ReceiverID: chatPartnerID,
		SenderID:   currentUserID,
		Content:    "Hello!",
	}

	response := message.ConvertMessageToResponseDTO(currentUserID)

	if response.ID != message.ID {
		t.Errorf("Expected ID %d, but got %d", message.ID, response.ID)
	}

	if response.CreatedAt != message.CreatedAt {
		t.Errorf("Expected created at %s, but got %s", message.CreatedAt, response.CreatedAt)
	}

	if response.IsSender != true {
		t.Errorf("Expected IsSender true, but got false")
	}

	if response.Content != message.Content {
		t.Errorf("Expected content %s, but got %s", message.Content, response.Content)
	}
}

func TestConvertMessagesToResponseDTO(t *testing.T) {
	currentUser := models.User{
		ID:        1,
		Username:  "currentuser",
		BirthDate: time.Now(),
		Name:      "John Doe",
		Bio:       "Lorem ipsum",
		Image:     nil,
	}

	chatPartner := models.User{
		ID:        2,
		Username:  "chatpartner",
		BirthDate: time.Now(),
		Name:      "Jane Smith",
		Bio:       "Dolor sit amet",
		Image:     nil,
	}

	messages := []models.Message{
		{
			ReceiverID: chatPartner.ID,
			SenderID:   currentUser.ID,
			Content:    "Hello!",
		},
		{
			ReceiverID: currentUser.ID,
			SenderID:   chatPartner.ID,
			Content:    "Hi!",
		},
	}

	response := models.ConvertMessagesToResponseDTO(messages, currentUser, chatPartner)

	if response.CurrentUser.ID != currentUser.ID {
		t.Errorf("Expected current user ID %d, but got %d", currentUser.ID, response.CurrentUser.ID)
	}

	if response.ChatPartner.ID != chatPartner.ID {
		t.Errorf("Expected chat partner ID %d, but got %d", chatPartner.ID, response.ChatPartner.ID)
	}

	if len(response.Messages) != len(messages) {
		t.Errorf("Expected %d messages, but got %d", len(messages), len(response.Messages))
	}
}
