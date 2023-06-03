package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID;not null"`
	PostID  uint   `gorm:"not null"`
	Comment string `gorm:"size:280" validate:"required,max=280"`
}

type CommentCreateRequestDTO struct {
	PostID  uint   `json:"postId" validate:"required"`
	Comment string `json:"comment" validate:"required,max=280"`
}
