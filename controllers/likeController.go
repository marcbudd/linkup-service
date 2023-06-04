package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/services"
)

func CreateLike(c *gin.Context) {

	// Get post id from url
	postID := c.Param("postID")
	if postID == "0" {
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

	// Create like
	err := services.CreateLike(userID.(uint), postID)

	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, gin.H{})

}

func DeleteLike(c *gin.Context) {

	// Get post id from url
	postID := c.Param("postID")
	if postID == "0" {
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

	// Delete like
	err := services.DeleteLike(userID.(uint), postID)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

func GetLikesByPostId(c *gin.Context) {

	// Get post id from url
	postID := c.Param("postID")
	if postID == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get likes
	likes, err := services.GetLikesByPostID(postID)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	//Respond
	c.JSON(http.StatusOK, likes)
}
