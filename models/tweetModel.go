package models

import (
	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	UserID  uint
	Content string `gorm:"size:280" validate:"required,max=280"`
}
