package services

import (
	"errors"
	"net/http"
	"sort"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/linkuperrors"
	"github.com/marcbudd/linkup-service/models"
)

func CreatePost(userID uint, req models.PostCreateRequestDTO) (*models.PostGetResponseDTO, *linkuperrors.LinkupError) {

	// Validate content
	if len(req.Content) > 280 {
		return nil, linkuperrors.New("content is over 280 chars", http.StatusBadRequest)
	}

	// Create post
	db := initalizers.DB
	post := *models.ConvertRequestDTOToPost(req, userID)

	result := db.Create(&post)
	if result.Error != nil {
		return nil, linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	// Post was just created, no likes, messages and is not liked by user
	numberLikes := int64(0)
	numberComments := int64(0)
	isLiked := false

	db.Preload("User").First(&post)

	return post.ConvertPostToResponseDTO(numberLikes, numberComments, isLiked), nil

}

func DeletePost(userID uint, postID string) *linkuperrors.LinkupError {

	// Get post
	db := initalizers.DB
	var post models.Post
	result := db.Where("id = ?", postID).First(&post)
	if result.Error != nil {
		return linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	// Check if user is owner of post
	if post.UserID != userID {
		return linkuperrors.New(errors.New("user is not owner of post").Error(), http.StatusForbidden)
	}

	// Delete likes and comments
	db.Where("post_id = ?", post.ID).Delete(&models.Like{})
	db.Where("post_id = ?", post.ID).Delete(&models.Comment{})

	// Delete post
	result = db.Delete(&post)
	if result.Error != nil {
		return linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func GetPostsByUserID(userID string, limit int, page int, currentUserID uint) ([]*models.PostGetResponseDTO, *linkuperrors.LinkupError) {
	// Set default values: Pagination
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	// Get posts
	db := initalizers.DB
	var posts []models.Post
	result := db.Where("user_id = ?", userID).Order("created_at DESC").Offset(offset).Limit(limit).Find(&posts)

	if result.Error != nil {
		return nil, linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	// Sort by created at desc
	sortByCreatedAtDesc(posts)

	// Create response dtos
	var dtos []*models.PostGetResponseDTO
	for _, post := range posts {
		db.Preload("User").First(&post)
		var numberLikes int64
		var numberComments int64
		db.Model(&models.Like{}).Where("post_id = ?", post.ID).Count(&numberLikes)
		db.Model(&models.Comment{}).Where("post_id = ?", post.ID).Count(&numberComments)

		// Check if current user is liking
		var count int64 = 0
		db.Model(&models.Like{}).Where("user_id = ? AND post_id = ?", currentUserID, post.ID).Count(&count)
		isLiked := count > 0

		dto := *post.ConvertPostToResponseDTO(numberLikes, numberComments, isLiked)
		dtos = append(dtos, &dto)
	}

	return dtos, nil
}

func GetPostByID(postID string, currentUserID uint) (*models.PostGetResponseDTO, *linkuperrors.LinkupError) {

	// Get post by id
	db := initalizers.DB
	var post models.Post
	result := db.Where("id = ?", postID).First(&post)
	if result.Error != nil {
		return nil, linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	// Create response dto
	db.Preload("User").First(&post)
	var numberLikes int64
	var numberComments int64
	db.Model(&models.Like{}).Where("post_id = ?", post.ID).Count(&numberLikes)
	db.Model(&models.Comment{}).Where("post_id = ?", post.ID).Count(&numberComments)

	// Check if current user is liking
	var count int64 = 0
	db.Model(&models.Like{}).Where("user_id = ? AND post_id = ?", currentUserID, post.ID).Count(&count)
	isLiked := count > 0

	responsePost := *post.ConvertPostToResponseDTO(numberLikes, numberComments, isLiked)

	return &responsePost, nil
}

func GetPostsForCurrentUser(currentUserID uint, limit int, page int) ([]*models.PostGetResponseDTO, *linkuperrors.LinkupError) {
	// Set default values: Pagination
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	// Get posts where user is following
	db := initalizers.DB
	var posts []models.Post

	result := db.
		Joins("JOIN follows ON follows.user_followed_id = posts.user_id").
		Where("follows.user_following_id = ? OR posts.user_id = ?", currentUserID, currentUserID).
		Distinct("posts.id").
		Order("posts.created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts)

	if result.Error != nil {
		return nil, linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	// Sort by created at desc
	sortByCreatedAtDesc(posts)

	var dtos []*models.PostGetResponseDTO
	for _, post := range posts {
		db.Preload("User").First(&post)
		var numberLikes int64
		var numberComments int64
		db.Model(&models.Like{}).Where("post_id = ?", post.ID).Count(&numberLikes)
		db.Model(&models.Comment{}).Where("post_id = ?", post.ID).Count(&numberComments)

		// Check if current user is liking
		var count int64 = 0
		db.Model(&models.Like{}).Where("user_id = ? AND post_id = ?", currentUserID, post.ID).Count(&count)
		isLiked := count > 0

		dto := *post.ConvertPostToResponseDTO(numberLikes, numberComments, isLiked)
		dtos = append(dtos, &dto)
	}

	return dtos, nil

}

func GetAllPosts(limit int, page int, currentUserID uint) ([]*models.PostGetResponseDTO, *linkuperrors.LinkupError) {
	// Set default values: Pagination
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	// Get posts where user is following
	db := initalizers.DB
	var posts []models.Post

	result := db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&posts)

	if result.Error != nil {
		return nil, linkuperrors.New(result.Error.Error(), http.StatusInternalServerError)
	}

	// Sort by created at desc
	sortByCreatedAtDesc(posts)

	// Create response dtos
	var dtos []*models.PostGetResponseDTO
	for _, post := range posts {
		db.Preload("User").First(&post)
		var numberLikes int64
		var numberComments int64
		db.Model(&models.Like{}).Where("post_id = ?", post.ID).Count(&numberLikes)
		db.Model(&models.Comment{}).Where("post_id = ?", post.ID).Count(&numberComments)

		// Check if current user is liking
		var count int64 = 0
		db.Model(&models.Like{}).Where("user_id = ? AND post_id = ?", currentUserID, post.ID).Count(&count)
		isLiked := count > 0

		dto := *post.ConvertPostToResponseDTO(numberLikes, numberComments, isLiked)
		dtos = append(dtos, &dto)
	}

	return dtos, nil

}

// Sort descending by created at attribute
func sortByCreatedAtDesc(posts []models.Post) {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})
}
