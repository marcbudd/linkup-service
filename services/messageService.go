package services

import (
	"errors"
	"sort"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"
)

func CreateMessage(userID uint, req models.MessageCreateRequestDTO) error {

	// Validate content
	if len(req.Content) > 280 {
		return errors.New("content is over 280 chars")
	}

	// Create message
	db := initalizers.DB
	message := models.Message{
		SenderID:   userID,
		ReceiverID: req.ReceiverID,
		Content:    req.Content,
	}

	result := db.Create(&message)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func GetMessagesByChat(userID uint, chatPartnerID string) (*models.MessagesOfChatGetResponseDTO, error) {

	// Get messages
	db := initalizers.DB
	var messages []models.Message

	result := db.Where("sender_id = ? OR sender_id = ?", userID, chatPartnerID).Where("receiver_id = ? OR receiver_id = ?", userID, chatPartnerID)

	if result.Error != nil {
		return nil, result.Error
	}

	// Find users
	currentUser := models.User{}
	chatPartner := models.User{}
	resultCurrentUser := db.Where("id = ?", userID).First(&currentUser)
	if resultCurrentUser.Error != nil {
		return nil, resultCurrentUser.Error
	}
	resultChatPartner := db.Where("id = ?", chatPartnerID).First(&chatPartner)
	if resultChatPartner.Error != nil {
		return nil, resultChatPartner.Error
	}

	// Create DTO
	messagesDTO := []models.MessageGetResponseDTO{}
	for _, message := range messages {
		messageDTO := models.MessageGetResponseDTO{
			ID:           message.ID,
			CreatedAt:    message.CreatedAt,
			ReceivedBool: message.ReceiverID == userID,
			Content:      message.Content,
		}

		messagesDTO = append(messagesDTO, messageDTO)
	}

	// Sort messages descending by date
	sort.Slice(messagesDTO, func(i, j int) bool {
		return messagesDTO[i].CreatedAt.After(messagesDTO[j].CreatedAt)
	})

	// Create complete response DTO
	responseDTO := models.MessagesOfChatGetResponseDTO{
		CurrentUser: models.UserGetResponseDTO{
			ID:        currentUser.ID,
			Username:  currentUser.Username,
			BirthDate: currentUser.BirthDate,
			Name:      currentUser.Name,
			Bio:       currentUser.Bio,
			Image:     currentUser.Image,
		},
		ChatPartner: models.UserGetResponseDTO{
			ID:        chatPartner.ID,
			Username:  chatPartner.Username,
			BirthDate: chatPartner.BirthDate,
			Name:      chatPartner.Name,
			Bio:       chatPartner.Bio,
			Image:     chatPartner.Image,
		},
		Messages: messagesDTO,
	}

	return &responseDTO, nil

}
