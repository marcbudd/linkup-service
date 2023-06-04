package services

import (
	"net/http"
	"sort"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/linkuperrors"
	"github.com/marcbudd/linkup-service/models"
)

func CreateMessage(currentUserID uint, req models.MessageCreateRequestDTO) *linkuperrors.LinkupError {

	// Validate content
	if len(req.Content) > 280 {
		return linkuperrors.New("content is over 280 chars", http.StatusBadRequest)
	}

	// Create message
	db := initalizers.DB
	message := models.ConvertRequestDTOToMessage(req, currentUserID)

	result := db.Create(&message)
	if result.Error != nil {
		return linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil

}

func GetMessagesByChat(currentUserID uint, chatPartnerID string) (*models.MessagesOfChatGetResponseDTO, *linkuperrors.LinkupError) {

	// Get messages
	db := initalizers.DB
	var messages []models.Message

	result := db.Where("sender_id = ? OR sender_id = ?", currentUserID, chatPartnerID).Where("receiver_id = ? OR receiver_id = ?", currentUserID, chatPartnerID)

	if result.Error != nil {
		return nil, linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	// Find users
	currentUser := models.User{}
	chatPartner := models.User{}
	resultCurrentUser := db.Where("id = ?", currentUserID).First(&currentUser)
	if resultCurrentUser.Error != nil {
		return nil, linkuperrors.New(resultCurrentUser.Error.Error(), http.StatusInternalServerError)
	}
	resultChatPartner := db.Where("id = ?", chatPartnerID).First(&chatPartner)
	if resultChatPartner.Error != nil {
		return nil, linkuperrors.New(resultChatPartner.Error.Error(), http.StatusInternalServerError)
	}

	// Sort messages descending by date
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].CreatedAt.After(messages[j].CreatedAt)
	})

	// Create complete response DTO
	responseDTO := *models.ConvertMessagesToResponseDTO(messages, currentUser, chatPartner)

	return &responseDTO, nil

}
