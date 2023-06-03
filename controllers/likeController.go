package controllers

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"
)

func CreateLike(c *gin.Context) {

	// Get post id from url
	param := c.Param("postId")

	postIdRaw, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	postId := uint(postIdRaw)

	// Get user id of logged in user
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Create like
	like := models.Like{
		UserID: userID.(uint),
		PostID: postId,
	}
	result := initalizers.DB.Create(&like)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create like",
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, gin.H{})

}

func DeleteLike(c *gin.Context) {

	// Get user id of logged in user
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Get post id from url
	postId := c.Param("postId")
	if postId == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get like
	var like models.Like
	result := initalizers.DB.Where("userId = ?", userID).Where("postId = ?", postId).First(&like)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Check if like belongs to user
	if like.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{})
	}

	// Delete like
	result = initalizers.DB.Delete(&like)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

func GetLikesByPostId(c *gin.Context) {

	// Get post id from url
	var postId = c.Param("postId")
	if postId == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	var likes []models.Like
	result := initalizers.DB.Where("post_id = ?", postId).Find(&likes)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Sort likes by date
	sort.Slice(likes, func(i, j int) bool {
		return likes[i].CreatedAt.Before(likes[j].CreatedAt)
	})

	//Respond
	c.JSON(http.StatusOK, gin.H{
		"likes": likes,
	})
}
