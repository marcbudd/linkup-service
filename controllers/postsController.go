package controllers

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"
)

func CreatePost(c *gin.Context) {
	//Get content of post from body
	var postCreateRequestDTO models.PostCreateRequestDTO

	if c.Bind(&postCreateRequestDTO) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	if len(postCreateRequestDTO.Content) > 280 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Content is over 280 chars",
		})

		return
	}

	// Get user id of logged in user
	userId := getCurrentUserId(c)
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	// Create post
	post := models.Post{Content: postCreateRequestDTO.Content, UserID: userId}
	result := initalizers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create post",
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, gin.H{
		"userId":  post.UserID,
		"content": post.Content,
	})

}

func DeletePost(c *gin.Context) {

	//Get user id of logged in user
	userId := getCurrentUserId(c)
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	// Get post id from url
	postId := c.Param("postId")
	if postId == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get post
	var post models.Post
	result := initalizers.DB.Where("id = ?", postId).First(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Check if post belongs to user
	if post.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{})
	}

	// Delete post
	result = initalizers.DB.Delete(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

func GetPostsByUserId(c *gin.Context) {

	// Get user id from url
	userId := c.Param("userId")
	if userId == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get posts
	var posts []models.Post
	result := initalizers.DB.Where("user_id = ?", userId).Find(&posts)
	sortByCreatedAtDesc(posts)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func getCurrentUserId(c *gin.Context) uint {

	// Benutzer aus dem Gin-Kontext abrufen
	userIdRaw, exists := c.Get("userId")
	if !exists {
		// Benutzer nicht gefunden
		return 0
	}

	userId, ok := userIdRaw.(uint)
	if !ok {
		// user id konnte nicht zu uint konvertiert werden
		return 0
	}

	return userId

}

// Sort descending by created at attribute
func sortByCreatedAtDesc(posts []models.Post) {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})
}
