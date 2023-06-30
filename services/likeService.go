package services

import (
	"net/http"
	"strconv"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/linkuperrors"
	"github.com/marcbudd/linkup-service/models"
)

func CreateLike(userID uint, postID string) *linkuperrors.LinkupError {

	// Check if user is already liking
	db := initalizers.DB
	var count int64 = 0
	db.Model(&models.Like{}).Where("user_id = ? AND post_id = ?", userID, postID).Count(&count)
	if count > 0 {
		return nil
	}

	// Convert string to uint
	temp, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		return linkuperrors.New(err.Error(), http.StatusBadRequest)
	}
	postIDUint := uint(temp)

	// Create like
	like := models.Like{
		UserID: userID,
		PostID: postIDUint,
	}
	result := db.Create(&like)

	if result.Error != nil {
		return linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func DeleteLike(userID uint, postID string) *linkuperrors.LinkupError {

	// Get like (if exists, else return nil)
	// Delete every like (if there are multiple)
	db := initalizers.DB
	var likes []models.Like
	result := db.Where("user_id = ? AND post_id = ?", userID, postID).Find(&likes)

	if result.Error != nil {
		return linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	// Delete likes
	for _, like := range likes {
		db.Delete(&like)
	}

	return nil
}

// Get likes of post
func GetLikesByPostID(postID string) ([]*models.LikeGetResponseDTO, *linkuperrors.LinkupError) {

	// Get likes
	db := initalizers.DB
	var likes []*models.Like
	result := db.Where("post_id = ?", postID).Find(&likes)

	if result.Error != nil {
		return nil, linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	// Create response object
	var dtos []*models.LikeGetResponseDTO
	for _, like := range likes {
		db.Preload("User").First(&like)
		dto := *like.ConvertLikeToResponseDTO()
		dtos = append(dtos, &dto)
	}

	return dtos, nil
}
