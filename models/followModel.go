package models

import (
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	UserFollowingID uint `gorm:"not null"`                            // User ID that is following
	UserFollowing   User `gorm:"foreignKey:UserFollowingID;not null"` // User that is following
	UserFollowedID  uint `gorm:"not null"`                            // User ID that is being followed
	UserFollowed    User `gorm:"foreignKey:UserFollowedID;not null"`  // User that is being followed
}

type FollowGetResponseDTO struct {
	ID       uint    `json:"id"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Image    *string `json:"image"`
}

// Function to convert follow to response dto
// can be called everywhere, changes can be made in one place
func (f *Follow) ConvertFollowingToResponseDTO() *FollowGetResponseDTO {
	return &FollowGetResponseDTO{
		ID:       f.UserFollowedID,
		Username: f.UserFollowed.Username,
		Name:     f.UserFollowed.Name,
		Image:    f.UserFollowed.Image,
	}
}

// Function to convert follow to response dto
// can be called everywhere, changes can be made in one place
func (f *Follow) ConvertFollowerToResponseDTO() *FollowGetResponseDTO {
	return &FollowGetResponseDTO{
		ID:       f.UserFollowingID,
		Username: f.UserFollowing.Username,
		Name:     f.UserFollowing.Name,
		Image:    f.UserFollowing.Image,
	}
}
