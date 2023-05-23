package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Tweets   []Tweet
}

// DTO struct für desialisierung von user jsons
type UserDTO struct {
	ID        uint `json:"ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Email     string
	Password  string
	Tweets    []Tweet
}
