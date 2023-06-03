package models

import (
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	UserFollowingID uint `gorm:"not null"`                   // User ID that is following
	UserFollowing   User `gorm:"foreignKey:UserID;not null"` // User that is following
	UserFollowedID  uint `gorm:"not null"`                   // User ID that is being followed
	UserFollowed    User `gorm:"foreignKey:UserID;not null"` // User that is being followed
}

type FollowingsOfUserGetResponseDTO struct {
	ID       uint    `json:"id"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Image    *string `json:"image"`
}

type FollowerOfUserGetResponseDTO struct {
	ID       uint    `json:"id"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Image    *string `json:"image"`
}
