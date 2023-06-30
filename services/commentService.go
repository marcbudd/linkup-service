package services

import (
	"net/http"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/linkuperrors"
	"github.com/marcbudd/linkup-service/models"
)

func CreateComment(userID uint, req models.CommentCreateRequestDTO) (*models.CommentGetResponseDTO, *linkuperrors.LinkupError) {

	// Validate content
	if len(req.Comment) > 280 {
		return nil, linkuperrors.New("comment is over 280 characters", http.StatusBadRequest)
	}

	// Create comment
	db := initalizers.DB
	comment := models.ConvertRequestDTOToComment(req, userID)

	result := db.Create(&comment)
	if result.Error != nil {
		return nil, linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	return comment.ConvertCommentToResponseDTO(), nil

}

func DeleteComment(userID uint, commentID string) *linkuperrors.LinkupError {

	// Get comment
	db := initalizers.DB
	var comment models.Comment

	result := db.Where("id = ?", commentID).First(&comment)
	if result.Error != nil {
		return linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	// Check if user is owner of comment
	if comment.UserID != userID {
		return linkuperrors.New("user is not owner of comment", http.StatusForbidden)
	}

	// Delete comment
	result = db.Delete(&comment)
	if result.Error != nil {
		return linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func GetCommentsByPostID(postID string) ([]*models.CommentGetResponseDTO, *linkuperrors.LinkupError) {

	// Get comments
	db := initalizers.DB
	var comments []*models.Comment
	result := db.Where("post_id = ?", postID).Find(&comments)

	if result.Error != nil {
		return nil, linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)

	}

	// converto to DTO
	var dtos []*models.CommentGetResponseDTO
	for _, comment := range comments {
		db.Preload("User").First(&comment)

		dto := *comment.ConvertCommentToResponseDTO()

		dtos = append(dtos, &dto)
	}

	return dtos, nil
}
