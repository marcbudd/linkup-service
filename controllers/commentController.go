package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/services"
)

// CreateComment creates a new comment.
// @Summary Create a comment
// @Description Create a new comment
// @Tags Comments
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param commentCreateRequestDTO body models.CommentCreateRequestDTO true "Comment creation data"
// @Success 201 {object} models.CommentGetResponseDTO
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /comments/{postID} [post]
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
	comment, err := services.CreateComment(userID.(uint), commentCreateRequestDTO)

	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, comment)

}

// DeleteComment deletes a comment.
// @Summary Delete a comment
// @Description Deletes a comment with the specified commentID
// @Tags Comments
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param commentID path string true "Comment ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /comments/{commentID} [delete]
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
	if commentID == "" {
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

// GetCommentsByPostID retrieves comments by post ID.
// @Summary Get comments by post ID
// @Description Retrieves comments associated with the specified postID
// @Tags Comments
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param postID path string true "Post ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /comments/posts/{postID} [get]
func GetCommentsByPostID(c *gin.Context) {

	// Get post id from url
	var postID = c.Param("postID")
	if postID == "" {
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
