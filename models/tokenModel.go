package models

import "time"

type Token struct {
	ID               uint      `gorm:"primary_key"`
	ExpirationDate   time.Time `gorm:"not null"`
	ConfirmationDate time.Time `gorm:"default:null"`
	TokenString      string    `gorm:"not null"`
	UserID           uint      `gorm:"not null"`
	User             User      `gorm:"foreignKey:UserID"`
}
