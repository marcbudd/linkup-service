package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID;not null"`
	Content string `gorm:"size:280;not null" validate:"required,max=280"`
}

type PostCreateRequestDTO struct {
	Content string `json:"content" validate:"required,max=280"`
}

// function to convert request dto to post
// can be called everywhere, changes can be made in one place
func ConvertRequestDTOToPost(req PostCreateRequestDTO, userID uint) *Post {
	return &Post{
		UserID:  userID,
		Content: req.Content,
	}
}

type PostGetResponseDTO struct {
	ID                 uint                `json:"id"`
	CreatedAt          time.Time           `json:"createdAt"`
	User               *UserGetResponseDTO `json:"user"`
	Content            string              `json:"content"`
	NumberLikes        int64               `json:"numberOfLikes"`
	NumberComments     int64               `json:"numberOfComments"`
	LikedByCurrentUser bool                `json:"likedByCurrentUser"`
}

// function to convert post to response dto
// can be called everywhere, changes can be made in one place
func (p *Post) ConvertPostToResponseDTO(numberLikes int64, numberComments int64, likedByCurrentUser bool) *PostGetResponseDTO {
	return &PostGetResponseDTO{
		ID:                 p.ID,
		CreatedAt:          p.CreatedAt,
		User:               p.User.ConvertUserToResponseDTO(),
		Content:            p.Content,
		NumberLikes:        numberLikes,
		NumberComments:     numberComments,
		LikedByCurrentUser: likedByCurrentUser,
	}
}
