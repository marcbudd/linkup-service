package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID;not null"`
	Content string `gorm:"size:280;not null" validate:"required,max=280"`
}

type PostCreateRequestDTO struct {
	Content string `json:"content" validate:"required,max=280"`
}

type PostGetResponseDTO struct {
	ID        uint               `json:"id"`
	CreatedAt string             `json:"createdAt"`
	User      UserGetResponseDTO `json:"user"`
	Content   string             `json:"content"`
}
