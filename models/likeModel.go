package models

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	UserID uint
	PostID uint
}
