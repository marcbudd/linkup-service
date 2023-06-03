package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint
	PostID  uint
	Comment string `gorm:"size:280" validate:"required,max=280"`
}

type CommentCreateRequestDTO struct {
	PostID  uint   `json:"postId" validate:"required"`
	Comment string `json:"comment" validate:"required,max=280"`
}
