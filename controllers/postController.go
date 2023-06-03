package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/services"
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

	// Get user id of logged in user
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Create post
	post, err := services.CreatePost(userID.(uint), postCreateRequestDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	// Respond
	c.JSON(http.StatusCreated, post)
}

func DeletePost(c *gin.Context) {

	// Get user id of logged in user
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Get post id from url
	postID := c.Param("postID")
	if postID == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Delete post
	err := services.DeletePost(userID.(uint), postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

func GetPostByPostID(c *gin.Context) {

	// Get post id from url
	postID := c.Param("postID")
	if postID == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get post
	post, err := services.GetPostByPostID(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	// Respond
	c.JSON(http.StatusOK, post)
}

func GetPostsByCurrentUser(c *gin.Context) {

	// Get query paramters
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 0
	}
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		page = 0
	}

	// Get user id of logged in user
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Get posts
	posts, err := services.GetPostsByUserID(string(userID.(uint)), int(limit), int(page))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	// Respond
	c.JSON(http.StatusOK, posts)
}

func GetPostsByUserID(c *gin.Context) {

	// Get query paramters
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 0
	}
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		page = 0
	}

	// Get user id from url
	userID := c.Param("userID")
	if userID == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get posts
	posts, err := services.GetPostsByUserID(userID, int(limit), int(page))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	// Respond
	c.JSON(http.StatusOK, posts)
}
