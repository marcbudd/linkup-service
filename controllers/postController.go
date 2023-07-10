package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/services"
)

// CreatePost creates a new post.
// @Summary Create a post
// @Description Creates a new post
// @Tags Posts
// @Accept json
// @Produce json
// @Param postCreateRequestDTO body models.PostCreateRequestDTO true "Post data"
// @Success 201 {object} models.PostGetResponseDTO
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/posts [post]
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
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, post)
}

// DeletePost deletes a post.
// @Summary Delete a post
// @Description Deletes a post
// @Tags Posts
// @Param postID path string true "Post ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /api/posts/{postID} [delete]
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
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Delete post
	err := services.DeletePost(userID.(uint), postID)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

// GetPostByID retrieves a specific post by its ID.
// @Summary Get a post by ID
// @Description Retrieves a specific post by its ID
// @Tags Posts
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200 {object} models.PostGetResponseDTO
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/posts/{postID} [get]
func GetPostByID(c *gin.Context) {

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

	// Get post
	post, err := services.GetPostByID(postID, userID.(uint))
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, post)
}

// GetPostsByUserID retrieves all posts of a specific user.
// @Summary Get posts by user ID
// @Description Retrieves all posts of a specific user
// @Tags Posts
// @Produce json
// @Param userID path string true "User ID"
// @Param limit query int false "Limit" default(0)
// @Param page query int false "Page" default(0)
// @Success 200 {array} models.PostGetResponseDTO
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/posts/user/{userID} [get]
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
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get user id of logged in user
	currentUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Get posts
	posts, serviceErr := services.GetPostsByUserID(userID, int(limit), int(page), currentUserID.(uint))
	if serviceErr != nil {
		c.JSON(serviceErr.HTTPStatusCode(), gin.H{
			"error": serviceErr.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, posts)
}

// GetPostsForCurrentUser returns posts of all users the logged in user follows
// @Summary Get posts for current user
// @Description Get all posts of the users the logged in user follows (feed)
// @Tags Posts
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Param limit query int false "Number of posts per page (default: 0 - all posts)"
// @Param page query int false "Page number (default: 0 - first page)"
// @Success 200 {array} models.PostGetResponseDTO
// @Failure 401
// @Failure 500
// @Router /api/posts/feed [get]
func GetPostsForCurrentUser(c *gin.Context) {
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
	posts, serviceErr := services.GetPostsForCurrentUser(userID.(uint), int(limit), int(page), userID.(uint))
	if serviceErr != nil {
		c.JSON(serviceErr.HTTPStatusCode(), gin.H{
			"error": serviceErr.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, posts)
}

// GetPosts returns all posts
// @Summary Get all posts
// @Description Get all posts in the system
// @Tags Posts
// @Produce json
// @Param limit query int false "Number of posts per page (default: 0 - all posts)"
// @Param page query int false "Page number (default: 0 - first page)"
// @Success 200 {array} models.PostGetResponseDTO
// @Failure 400
// @Failure 500
// @Router /api/posts [get]
func GetPosts(c *gin.Context) {
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
	posts, serviceErr := services.GetAllPosts(int(limit), int(page), userID.(uint))
	if serviceErr != nil {
		c.JSON(serviceErr.HTTPStatusCode(), gin.H{
			"error": serviceErr.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, posts)
}

// CreatePostGPT returns a post generated by OPENAI ChatGPT
// @Summary Create post using ChatGPT
// @Description Create a post generated by OPENAI ChatGPT
// @Tags Posts
// @Produce json
// @Success 200 {array} string
// @Failure 500
// @Router /api/posts/gpt [get]
func CreatePostGPT(c *gin.Context) {
	// Get string from GPT
	string, err := services.CreatePostGPT()
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, string)
}
