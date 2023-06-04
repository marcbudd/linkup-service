package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/services"
)

func CreateComment(c *gin.Context) {

	//Read body
	var commentCreateRequestDTO models.CommentCreateRequestDTO
	if c.Bind(&commentCreateRequestDTO) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
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

	// Create comment
	err := services.CreateComment(userID.(uint), commentCreateRequestDTO)

	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, gin.H{})

}

func DeleteComment(c *gin.Context) {

	// Get user id of logged in user
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Get comment id from url
	commentID := c.Param("commentID")
	if commentID == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Delete comment
	err := services.DeleteComment(userID.(uint), commentID)

	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

func GetCommentsByPostId(c *gin.Context) {

	// Get post id from url
	var postID = c.Param("postID")
	if postID == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	comments, err := services.GetCommentsByPostID(postID)

	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	//Respond
	c.JSON(http.StatusOK, comments)
}
