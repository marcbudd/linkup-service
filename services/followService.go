package services

import (
	"net/http"
	"strconv"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/linkuperrors"
	"github.com/marcbudd/linkup-service/models"
)

func CreateFollow(userIDFollowing uint, userIDFollowed string) *linkuperrors.LinkupError {

	// Check if user is already following
	db := initalizers.DB
	var count int64 = 0
	db.Model(&models.Follow{}).Where("user_following_id = ? AND user_followed_id = ?", userIDFollowing, userIDFollowed).Count(&count)
	if count > 0 {
		return nil
	}

	// Check if user wants to follow himself
	if strconv.Itoa(int(userIDFollowing)) == userIDFollowed {
		return linkuperrors.New("user cannot follow himself", http.StatusBadRequest)
	}

	// Convert string to uint
	temp, err := strconv.ParseUint(userIDFollowed, 10, 64)
	if err != nil {
		return linkuperrors.New("invalid user id", http.StatusBadRequest)
	}
	userIDFollowedUint := uint(temp)

	// Create follow
	follow := models.Follow{
		UserFollowingID: userIDFollowing,
		UserFollowedID:  userIDFollowedUint,
	}
	result := db.Create(&follow)

	if result.Error != nil {
		return linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func DeleteFollow(userIDFollowing uint, userIDFollowed string) *linkuperrors.LinkupError {

	// Get follow (if exists, else return nil)
	// Delete every follow (if there are multiple)
	db := initalizers.DB
	var follows []models.Follow
	result := db.Where("user_following_id = ? AND user_followed_id = ?", userIDFollowing, userIDFollowed).Find(&follows)

	if result.Error != nil {
		return linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)

	}

	// Delete likes
	for _, follow := range follows {
		db.Delete(&follow)
	}

	return nil
}

// Get lists of user that are followed by a user
func GetFollowingsByUserID(userID string) ([]*models.FollowGetResponseDTO, *linkuperrors.LinkupError) {

	// Get followings of a user
	db := initalizers.DB
	var follows []*models.Follow
	result := db.Where("user_following_id = ?", userID).Preload("UserFollowing").Preload("UserFollowed").Find(&follows)
	if result.Error != nil {
		return nil, linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)

	}

	// Create response object
	var dtos []*models.FollowGetResponseDTO
	for _, follow := range follows {

		dto := *follow.ConvertFollowingToResponseDTO()

		dtos = append(dtos, &dto)
	}

	return dtos, nil
}

// Get lists of user that are following user
func GetFollowersByUserID(userID string) ([]*models.FollowGetResponseDTO, *linkuperrors.LinkupError) {

	// Get follower of a user
	db := initalizers.DB
	var follows []*models.Follow
	result := db.Where("user_followed_id = ?", userID).Preload("UserFollowing").Preload("UserFollowed").Find(&follows)
	if result.Error != nil {
		return nil, linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)

	}

	// Create response object
	var dtos []*models.FollowGetResponseDTO
	for _, follow := range follows {
		dto := *follow.ConvertFollowerToResponseDTO()

		dtos = append(dtos, &dto)
	}

	return dtos, nil
}
