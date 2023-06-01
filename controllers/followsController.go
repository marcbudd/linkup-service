package controllers

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"
)

func CreateFollow(c *gin.Context) {

	// Get content from body
	var followCreateRequestDTO models.FollowCreateRequestDTO

	if c.Bind(&followCreateRequestDTO) != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	//Get user id of logged in user
	userId := getCurrentUserId(c)
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	// Create follow
	follow := models.Follow{
		FollowingID:  followCreateRequestDTO.UserID,
		FollowedByID: userId,
	}
	result := initalizers.DB.Create(&follow)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create follow",
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, gin.H{})

}

func DeleteFollow(c *gin.Context) {

	//Get user id of logged in user
	userId := getCurrentUserId(c)
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	//TODO: change here to get follow id from url
	// Get user id from url
	followId := c.Param("followId")
	if followId == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get follow
	var follow models.Follow
	result := initalizers.DB.Where("id = ?", followId).First(&follow)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Check if follow belongs to user
	if follow.FollowedByID != userId {
		c.JSON(http.StatusForbidden, gin.H{})
	}

	// Delete follow
	result = initalizers.DB.Delete(&follow)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

func GetFollowsByUserId(c *gin.Context) {

	// Get user id from url
	var userId = c.Param("userId")
	if userId == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	var follows []models.Follow
	result := initalizers.DB.Where("followed_by_id = ?", userId).Find(&follows)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Sort follows by username
	sort.Slice(follows, func(i, j int) bool {
		return follows[i].Following.Username < follows[j].Following.Username
	})

	//Respond
	c.JSON(http.StatusOK, gin.H{
		"follows": follows,
	})
}

func GetFollowingsByUserId(c *gin.Context) {

	// Get user id from url
	var userId = c.Param("userId")
	if userId == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	var follows []models.Follow
	result := initalizers.DB.Where("following_id = ?", userId).Find(&follows)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Sort follows by username
	sort.Slice(follows, func(i, j int) bool {
		return follows[i].FollowedBy.Username < follows[j].FollowedBy.Username
	})

	//Respond
	c.JSON(http.StatusOK, gin.H{
		"follows": follows,
	})
}
