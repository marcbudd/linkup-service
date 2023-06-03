package services

import (
	"strconv"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"
)

func CreateLike(userID uint, postID string) error {

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
		return err
	}
	postIDUint := uint(temp)

	// Create like
	like := models.Like{
		UserID: userID,
		PostID: postIDUint,
	}
	result := db.Create(&like)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteLike(userID uint, postID string) error {

	// Get like (if exists, else return nil)
	// Delete every like (if there are multiple)
	db := initalizers.DB
	var likes []models.Like
	result := db.Where("user_id = ? AND post_id = ?", userID, postID).Find(&likes)

	if result.Error != nil {
		return result.Error
	}

	// Delete likes
	for _, like := range likes {
		db.Delete(&like)
	}

	return nil
}

// Get likes of post
func GetLikesByPostID(postID string) ([]*models.LikesOfPostGetResponseDTO, error) {

	// Get likes
	db := initalizers.DB
	var likes []*models.Like
	result := db.Where("post_id = ?", postID).Find(&likes)

	if result.Error != nil {
		return nil, result.Error
	}

	// Create response object
	var dtos []*models.LikesOfPostGetResponseDTO
	for _, like := range likes {
		dto := models.LikesOfPostGetResponseDTO{
			ID:       like.User.ID,
			Username: like.User.Username,
			Name:     like.User.Name,
			Image:    like.User.Image,
		}

		dtos = append(dtos, &dto)
	}

	return dtos, nil
}
