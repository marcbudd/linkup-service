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

type MessageGetResponseDTO struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	ReceivedBool bool      `json:"receivedBool"` //is true if the message was written to you, false if you wrote it
	Content      string    `json:"text"`
}

type MessagesOfChatGetResponseDTO struct {
	CurrentUser UserGetResponseDTO      `json:"currentUser"` // the user that is logged in, making the request
	ChatPartner UserGetResponseDTO      `json:"chatPartner"` // the user that is the chat partner of the current chat
	Messages    []MessageGetResponseDTO `json:"messages"`
}
