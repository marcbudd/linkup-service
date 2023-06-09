package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/services"
)

// CreateLike creates a like for a post.
// @Summary Create a like for a post
// @Description Creates a like for the post with the specified postID
// @Tags Likes
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param postID path string true "Post ID"
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /likes/{postID} [post]
func CreateLike(c *gin.Context) {

	// Get post id from url
	postID := c.Param("postID")
	if postID == "" {
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

// DeleteLike deletes a like for a post.
// @Summary Delete a like for a post
// @Description Deletes the like for the post with the specified postID
// @Tags Likes
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param postID path string true "Post ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /likes/{postID} [delete]
func DeleteLike(c *gin.Context) {

	// Get post id from url
	postID := c.Param("postID")
	if postID == "" {
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

// GetLikesByPostId retrieves the likes for a post.
// @Summary Get likes for a post
// @Description Retrieves the likes for the post with the specified postID
// @Tags Likes
// @Param postID path string true "Post ID"
// @Success 200 {array} models.LikeGetResponseDTO
// @Failure 400
// @Failure 500
// @Router /likes/{postID} [get]
func GetLikesByPostId(c *gin.Context) {

	// Get post id from url
	postID := c.Param("postID")
	if postID == "" {
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
