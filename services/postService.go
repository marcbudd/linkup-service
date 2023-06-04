package services

import (
	"errors"
	"sort"

	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"
)

func CreatePost(userID uint, req models.PostCreateRequestDTO) (*models.Post, error) {

	// Validate content
	if len(req.Content) > 280 {
		return nil, errors.New("content is over 280 chars")
	}

	// Create post
	db := initalizers.DB
	post := *models.ConvertRequestDTOToPost(req, userID)

	result := db.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil

}

func DeletePost(userID uint, postID string) error {

	// Get post
	db := initalizers.DB
	var post models.Post
	result := db.Where("id = ?", postID).First(&post)
	if result.Error != nil {
		return result.Error
	}

	// Check if user is owner of post
	if post.UserID != userID {
		return errors.New("user is not owner of post")
	}

	// Delete post
	result = db.Delete(&post)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetPostsByUserID(userID string, limit int, page int) ([]*models.PostGetResponseDTO, error) {
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
	result := db.Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	// Sort by created at desc
	sortByCreatedAtDesc(posts)

	var dtos []*models.PostGetResponseDTO
	for _, post := range posts {
		dto := *post.ConvertPostToResponseDTO()
		dtos = append(dtos, &dto)
	}

	return dtos, nil
}

func GetPostByID(postID string) (*models.PostGetResponseDTO, error) {

	// Get post by id
	db := initalizers.DB
	var post models.Post
	result := db.Where("id = ?", postID).First(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	// Create response dto
	responsePost := *post.ConvertPostToResponseDTO()

	return &responsePost, nil
}

func GetPostsForCurrentUser(userID string, limit int, page int) ([]*models.PostGetResponseDTO, error) {
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
		Where("follows.user_following_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	// Sort by created at desc
	sortByCreatedAtDesc(posts)

	var dtos []*models.PostGetResponseDTO
	for _, post := range posts {
		dto := *post.ConvertPostToResponseDTO()
		dtos = append(dtos, &dto)
	}

	return dtos, nil

}

func GetAllPosts(limit int, page int) ([]*models.PostGetResponseDTO, error) {
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

	result := db.Offset(offset).Limit(limit).Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	// Sort by created at desc
	sortByCreatedAtDesc(posts)

	// Create response dtos
	var dtos []*models.PostGetResponseDTO
	for _, post := range posts {
		dto := *post.ConvertPostToResponseDTO()
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
