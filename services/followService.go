package services

import (
	"strconv"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"
)

func CreateFollow(userIDFollowing uint, userIDFollowed string) error {

	// Check if user is already following
	db := initalizers.DB
	var count int64 = 0
	db.Model(&models.Follow{}).Where("user_following_id = ? AND user_followed_id = ?", userIDFollowing, userIDFollowed).Count(&count)
	if count > 0 {
		return nil
	}

	// Convert string to uint
	temp, err := strconv.ParseUint(userIDFollowed, 10, 64)
	if err != nil {
		return err
	}
	userIDFollowedUint := uint(temp)

	// Create follow
	follow := models.Follow{
		UserFollowingID: userIDFollowing,
		UserFollowedID:  userIDFollowedUint,
	}
	result := db.Create(&follow)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteFollow(userIDFollowing uint, userIDFollowed string) error {

	// Get follow (if exists, else return nil)
	// Delete every follow (if there are multiple)
	db := initalizers.DB
	var follows []models.Follow
	result := db.Where("user_following_id = ? AND user_followed_id = ?", userIDFollowing, userIDFollowed).Find(&follows)

	if result.Error != nil {
		return result.Error
	}

	// Delete likes
	for _, follow := range follows {
		db.Delete(&follow)
	}

	return nil
}

// Get lists of user that are followed by a user
func GetFollowingsOfUserID(userID string) ([]*models.FollowingsOfUserGetResponseDTO, error) {

	// Get follows of a user
	db := initalizers.DB
	var follows []*models.Follow
	result := db.Where("user_follwing_id = ?", userID).Find(&follows)
	if result.Error != nil {
		return nil, result.Error
	}

	// Create response object
	var dtos []*models.FollowingsOfUserGetResponseDTO
	for _, follow := range follows {
		dto := models.FollowingsOfUserGetResponseDTO{
			ID:       follow.UserFollowed.ID,
			Username: follow.UserFollowed.Username,
			Name:     follow.UserFollowed.Name,
			Image:    follow.UserFollowed.Image,
		}

		dtos = append(dtos, &dto)
	}

	return dtos, nil
}

// Get lists of user that are following user
func GetFollowerOfUserID(userID string) ([]*models.FollowerOfUserGetResponseDTO, error) {

	// Get follower of a user
	db := initalizers.DB
	var follows []*models.Follow
	result := db.Where("user_followed_id = ?", userID).Find(&follows)
	if result.Error != nil {
		return nil, result.Error
	}

	// Create response object
	var dtos []*models.FollowerOfUserGetResponseDTO
	for _, follow := range follows {
		dto := models.FollowerOfUserGetResponseDTO{
			ID:       follow.UserFollowing.ID,
			Username: follow.UserFollowing.Username,
			Name:     follow.UserFollowing.Name,
			Image:    follow.UserFollowing.Image,
		}

		dtos = append(dtos, &dto)
	}

	return dtos, nil
}
