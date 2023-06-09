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

type LikeGetResponseDTO struct {
	ID       uint    `json:"id"`
	UserID   uint    `json:"userID"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Image    *string `json:"image"`
}

// Function to convert like to response dto
// can be called everywhere, changes can be made in one place
func (l *Like) ConvertLikeToResponseDTO() *LikeGetResponseDTO {
	return &LikeGetResponseDTO{
		ID:       l.ID,
		UserID:   l.UserID,
		Username: l.User.Username,
		Name:     l.User.Name,
		Image:    l.User.Image,
	}
}
