package services

import (
	"errors"
	"sort"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"
)

func CreateMessage(currentUserID uint, req models.MessageCreateRequestDTO) error {

	// Validate content
	if len(req.Content) > 280 {
		return errors.New("content is over 280 chars")
	}

	// Create message
	db := initalizers.DB
	message := models.ConvertRequestDTOToMessage(req, currentUserID)

	result := db.Create(&message)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func GetMessagesByChat(currentUserID uint, chatPartnerID string) (*models.MessagesOfChatGetResponseDTO, error) {

	// Get messages
	db := initalizers.DB
	var messages []models.Message

	result := db.Where("sender_id = ? OR sender_id = ?", currentUserID, chatPartnerID).Where("receiver_id = ? OR receiver_id = ?", currentUserID, chatPartnerID)

	if result.Error != nil {
		return nil, result.Error
	}

	// Find users
	currentUser := models.User{}
	chatPartner := models.User{}
	resultCurrentUser := db.Where("id = ?", currentUserID).First(&currentUser)
	if resultCurrentUser.Error != nil {
		return nil, resultCurrentUser.Error
	}
	resultChatPartner := db.Where("id = ?", chatPartnerID).First(&chatPartner)
	if resultChatPartner.Error != nil {
		return nil, resultChatPartner.Error
	}

	// Sort messages descending by date
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].CreatedAt.After(messages[j].CreatedAt)
	})

	// Create complete response DTO
	responseDTO := *models.ConvertMessagesToResponseDTO(messages, currentUser, chatPartner)

	return &responseDTO, nil

}
