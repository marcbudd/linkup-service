package controllers

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"
)

func CreateComment(c *gin.Context) {
	//Get content of post from body
	var commentCreateRequestDTO models.CommentCreateRequestDTO

	if c.Bind(&commentCreateRequestDTO) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	if len(commentCreateRequestDTO.Comment) > 280 {
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

	// Create comment
	comment := models.Comment{
		UserID:  userId,
		PostID:  commentCreateRequestDTO.PostID,
		Comment: commentCreateRequestDTO.Comment,
	}
	result := initalizers.DB.Create(&comment)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create comment",
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, gin.H{
		"userId":  comment.UserID,
		"content": comment.Comment,
	})

}

func DeleteComment(c *gin.Context) {

	//Get user id of logged in user
	userId := getCurrentUserId(c)
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	// Get comment id from url
	commentId := c.Param("commentId")
	if commentId == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get comment
	var comment models.Comment
	result := initalizers.DB.Where("id = ?", commentId).First(&comment)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Check if comment belongs to user
	if comment.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{})
	}

	// Delete comment
	result = initalizers.DB.Delete(&comment)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

func GetCommentsByPostId(c *gin.Context) {

	// Get post id from url
	var postId = c.Param("postId")
	if postId == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	var comments []models.Comment
	result := initalizers.DB.Where("post_id = ?", postId).Find(&comments)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Sort comments by date
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].CreatedAt.Before(comments[j].CreatedAt)
	})

	//Respond
	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})
}
