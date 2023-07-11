package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/services"
)

// CreateFollow creates a follow relationship between users.
// @Summary Create a follow relationship
// @Description Creates a follow relationship between the logged-in user and the user with the specified followedUserID
// @Tags Follows
// @Param followedUserID path string true "Followed User ID"
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/follows/{followedUserID} [post]
func CreateFollow(c *gin.Context) {

	// Get user id from url
	var followedUserID = c.Param("followedUserID")
	if followedUserID == "" {
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

// DeleteFollow deletes a follow relationship between users.
// @Summary Delete a follow relationship
// @Description Deletes the follow relationship between the logged-in user and the user with the specified followedUserID
// @Tags Follows
// @Param followedUserID path string true "Followed User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/follows/{followedUserID} [delete]
func DeleteFollow(c *gin.Context) {

	// Get user id from url
	var followedUserID = c.Param("followedUserID")
	if followedUserID == "" {
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

// GetFollowingsOfUserID retrieves the followings of a user by userID.
// @Summary Get followings of a user
// @Description Retrieves the list of followings for the user with the specified userID
// @Tags Follows
// @Produce json
// @Param userID path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/follows/{userID}/followings [get]
func GetFollowingsByUserID(c *gin.Context) {

	// Get user id from url
	var userID = c.Param("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get followings of user
	followings, err := services.GetFollowingsByUserID(userID)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, followings)

}

// GetFollowersByUserID retrieves the followers of a user by userID.
// @Summary Get followers of a user
// @Description Retrieves the list of followers for the user with the specified userID
// @Tags Follows
// @Produce json
// @Param userID path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/follows/{userID}/followers [get]
func GetFollowersByUserID(c *gin.Context) {

	// Get user id from url
	var userID = c.Param("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get follower of a user
	followers, err := services.GetFollowersByUserID(userID)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, followers)

}
