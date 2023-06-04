package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID;not null"`
	PostID  uint   `gorm:"not null"`
	Comment string `gorm:"size:280" validate:"required,max=280"`
}

type CommentCreateRequestDTO struct {
	PostID  uint   `json:"postId" validate:"required"`
	Comment string `json:"comment" validate:"required,max=280"`
}

// function to convert comment to response dto
// can be called everywhere, changes can be made in one place
func ConvertRequestDTOToComment(req CommentCreateRequestDTO, userID uint) *Comment {
	return &Comment{
		UserID:  userID,
		PostID:  req.PostID,
		Comment: req.Comment,
	}
}

type CommentGetResponseDTO struct {
	ID        uint               `json:"id"`
	CreatedAt time.Time          `json:"createdAt"`
	User      UserGetResponseDTO `json:"user"`
	Comment   string             `json:"comment"`
}

// function to convert comment to respone dto
// can be called everywhere, changes can be made in one place
func (c *Comment) ConvertCommentToResponseDTO() *CommentGetResponseDTO {
	return &CommentGetResponseDTO{
		ID:        c.ID,
		CreatedAt: c.CreatedAt,
		User:      *c.User.ConvertUserToResponseDTO(),
		Comment:   c.Comment,
	}
}
