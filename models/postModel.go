package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID  uint
	Content string `gorm:"size:280" validate:"required,max=280"`
}

type PostCreateRequestDTO struct {
	Content string `json:"content" validate:"required,max=280"`
}
