package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/services"
)

func CreateFollow(c *gin.Context) {

	// Get user id from url
	var followedUserID = c.Param("followedUserID")
	if followedUserID == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get user id of logged in user
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Create follow
	err := services.CreateFollow(userID.(uint), followedUserID)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, gin.H{})
}

func DeleteFollow(c *gin.Context) {

	// Get user id from url
	var followedUserID = c.Param("followedUserID")
	if followedUserID == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get user id of logged in user
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Delete follow
	err := services.DeleteFollow(userID.(uint), followedUserID)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

func GetFollowingsOfUserID(c *gin.Context) {

	// Get user id from url
	var userID = c.Param("userID")
	if userID == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get followings of user
	followings, err := services.GetFollowingsOfUserID(userID)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, followings)

}

func GetFollowersOfUserID(c *gin.Context) {

	// Get user id from url
	var userID = c.Param("userID")
	if userID == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get follower of a user
	followers, err := services.GetFollowersOfUserID(userID)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, followers)

}
