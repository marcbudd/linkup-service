package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SenderID   uint   `gorm:"not null"`
	Sender     User   `gorm:"foreignKey:SenderID; not null"`
	ReceiverID uint   `gorm:"not null"`
	Receiver   User   `gorm:"foreignKey:ReceiverID; not null"`
	Content    string `gorm:"size:400; not null" validate:"required,max=280"`
}

type MessageCreateRequestDTO struct {
	ReceiverID uint   `json:"receiverID" validate:"required"`
	Content    string `json:"text" validate:"required, max=280"`
}

// function to convert request dto to message
// can be called everywhere, changes can be made in one place
func ConvertRequestDTOToMessage(req MessageCreateRequestDTO, senderID uint) *Message {
	return &Message{
		SenderID:   senderID,
		ReceiverID: req.ReceiverID,
		Content:    req.Content,
	}
}

type MessageGetResponseDTO struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	IsSender  bool      `json:"isSender"` //is true if the message was written by you, false if it was written to you
	Content   string    `json:"text"`
}

// function to convert message to response dto
// can be called everywhere, changes can be made in one place
func (m *Message) ConvertMessageToResponseDTO(currentUserID uint) *MessageGetResponseDTO {
	return &MessageGetResponseDTO{
		ID:        m.ID,
		CreatedAt: m.CreatedAt,
		IsSender:  m.SenderID == currentUserID,
		Content:   m.Content,
	}
}

type MessagesOfChatGetResponseDTO struct {
	CurrentUser UserGetResponseDTO      `json:"currentUser"` // the user that is logged in, making the request
	ChatPartner UserGetResponseDTO      `json:"chatPartner"` // the user that is the chat partner of the current chat
	Messages    []MessageGetResponseDTO `json:"messages"`
}

// function to convert messages to response dto
// can be called everywhere, changes can be made in one place
func ConvertMessagesToResponseDTO(messages []Message, currentUser User, chatPartner User) *MessagesOfChatGetResponseDTO {
	var messageDTOs []MessageGetResponseDTO
	for _, message := range messages {
		message.ConvertMessageToResponseDTO(currentUser.ID)
		messageDTOs = append(messageDTOs, *message.ConvertMessageToResponseDTO(message.SenderID))
	}

	return &MessagesOfChatGetResponseDTO{
		CurrentUser: *currentUser.ConvertUserToResponseDTO(),
		ChatPartner: *chatPartner.ConvertUserToResponseDTO(),
		Messages:    messageDTOs,
	}
}
