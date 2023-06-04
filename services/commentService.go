package services

import (
	"errors"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"
)

func CreateComment(userID uint, req models.CommentCreateRequestDTO) error {

	// Validate content
	if len(req.Comment) > 280 {
		return errors.New("content is over 280 chars")
	}

	// Create comment
	db := initalizers.DB
	comment := models.ConvertRequestDTOToComment(req, userID)

	result := db.Create(&comment)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func DeleteComment(userID uint, commentID string) error {

	// Get comment
	db := initalizers.DB
	var comment models.Comment

	result := db.Where("id = ?", commentID).First(&comment)
	if result.Error != nil {
		return result.Error
	}

	// Check if user is owner of comment
	if comment.UserID != userID {
		return errors.New("user is not owner of comment")
	}

	// Delete comment
	result = db.Delete(&comment)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetCommentsByPostID(postID string) ([]*models.CommentGetResponseDTO, error) {

	// Get comments
	db := initalizers.DB
	var comments []*models.Comment
	result := db.Where("post_id = ?", postID).Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	}

	// converto to DTO
	var dtos []*models.CommentGetResponseDTO
	for _, comment := range comments {
		dto := *comment.ConvertCommentToResponseDTO()

		dtos = append(dtos, &dto)
	}

	return dtos, nil
}
