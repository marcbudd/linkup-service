package models

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model `gorm:"not null"`
	UserID     uint
	User       User `gorm:"foreignKey:UserID;not null"`
	PostID     uint `gorm:"not null"`
	Post       Post `gorm:"foreignKey:PostID;not null"`
}

type LikesOfPostGetResponseDTO struct {
	ID       uint    `json:"id"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Image    *string `json:"image"`
}
