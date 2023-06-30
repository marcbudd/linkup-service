package services

import (
	"net/http"
	"sort"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/linkuperrors"
	"github.com/marcbudd/linkup-service/models"
)

func CreateMessage(currentUserID uint, req models.MessageCreateRequestDTO) (*models.MessageGetResponseDTO, *linkuperrors.LinkupError) {

	// Validate content
	if len(req.Content) > 280 {
		return nil, linkuperrors.New("content is over 280 chars", http.StatusBadRequest)
	}

	// Create message
	db := initalizers.DB
	message := models.ConvertRequestDTOToMessage(req, currentUserID)

	result := db.Create(&message)
	if result.Error != nil {
		return nil, linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	return message.ConvertMessageToResponseDTO(currentUserID), nil

}

func GetMessagesByChat(currentUserID uint, chatPartnerID string) (*models.MessagesOfChatGetResponseDTO, *linkuperrors.LinkupError) {

	// Get messages
	db := initalizers.DB
	var messages []models.Message

	result := db.Where("sender_id = ? OR sender_id = ?", currentUserID, chatPartnerID).Where("receiver_id = ? OR receiver_id = ?", currentUserID, chatPartnerID).Find(&messages)

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

	// Set IsReadyByReceiver to true, where ReceiverID = currentUserID
	for _, message := range messages {
		db.Preload("Sender").First(&message)
		db.Preload("Receiver").First(&message)

		if message.ReceiverID == currentUserID {
			message.IsReadByReceiver = true
			db.Save(&message)
		}
	}

	// Create complete response DTO
	responseDTO := *models.ConvertMessagesToResponseDTO(messages, currentUser, chatPartner)

	return &responseDTO, nil

}

func GetChatsByUserID(currentUserID uint) ([]*models.ChatOfUserGetResponseDTO, *linkuperrors.LinkupError) {
	var chats []*models.ChatOfUserGetResponseDTO

	db := initalizers.DB

	// Find all the chat partners of current user id
	var chatPartnerIDs []uint
	sql :=
		`
		SELECT DISTINCT id 
		FROM
		((SELECT DISTINCT receiver_id AS id
		FROM messages
		WHERE sender_id = ?)
		UNION 
		(SELECT DISTINCT sender_id AS id
		FROM messages
		WHERE receiver_id = ?))
		`
	err := db.Raw(sql, currentUserID, currentUserID).Scan(&chatPartnerIDs).Error
	if err != nil {
		return nil, linkuperrors.New(err.Error(), http.StatusInternalServerError)
	}

	// For each chat partner get user and LastMassage
	for _, chatPartnerID := range chatPartnerIDs {
		// Get chat partner data
		var user models.User
		err := db.Where("user_id = ?", chatPartnerID).Find(&user).Error
		if err != nil {
			return nil, linkuperrors.New(err.Error(), http.StatusInternalServerError)
		}

		// Get last message
		var message models.Message
		err = db.Where("sender_id = ? OR sender_id = ?", currentUserID, chatPartnerID).Where("receiver_id = ? OR receiver_id = ?", currentUserID, chatPartnerID).Order("created_at DESC").Limit(1).Find(&message).Error
		if err != nil {
			return nil, linkuperrors.New(err.Error(), http.StatusInternalServerError)
		}

		// Get number of unread messages
		var numberOfUnreadMessages int64
		err = db.Where("receiver_id = ? AND senderID = ? AND IsReadByReceiver = false", currentUserID, chatPartnerID).Count(&numberOfUnreadMessages).Error
		if err != nil {
			return nil, linkuperrors.New(err.Error(), http.StatusInternalServerError)
		}

		// Create dto and append to output array
		dto := models.ConvertChatsToResponseDTO(&message, &user, &numberOfUnreadMessages)
		chats = append(chats, dto)

	}

	return chats, nil

}
