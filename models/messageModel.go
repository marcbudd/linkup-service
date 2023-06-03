package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderID   uint
	Sender     User `gorm:"foreignKey:SenderID"`
	ReceiverID uint
	Receiver   User `gorm:"foreignKey:ReceiverID"`
	Text       string
}

type MessageCreateRequestDTO struct {
	ReceiverID uint   `json:"receiverID" validate:"required"`
	Text       string `json:"text" validate:"required"`
}
