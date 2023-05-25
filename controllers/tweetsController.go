package controllers

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"
)

func PostTweet(c *gin.Context) {
	//Get content of tweet from body
	var body struct {
		Content string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	if len(body.Content) > 280 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Content is over 280 chars",
		})

		return
	}

	// Get user id of logged in user
	userId := getCurrentUserId(c)
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	// Create tweet
	tweet := models.Tweet{Content: body.Content, UserID: userId}
	result := initalizers.DB.Create(&tweet)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create tweet",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"userId":  tweet.UserID,
		"content": tweet.Content,
	})

}

func DeleteTweet(c *gin.Context) {

	//Get user id of logged in user
	userId := getCurrentUserId(c)
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	// Get tweet id from url
	tweetId := c.Param("tweetId")
	if tweetId == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get tweet
	var tweet models.Tweet
	result := initalizers.DB.Where("id = ?", tweetId).First(&tweet)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Check if tweet belongs to user
	if tweet.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{})
	}

	// Delete tweet
	result = initalizers.DB.Delete(&tweet)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

func GetTweetsByUserId(c *gin.Context) {

	// Get user id from url
	userId := c.Param("userId")
	if userId == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get tweets
	var tweets []models.Tweet
	result := initalizers.DB.Where("user_id = ?", userId).Find(&tweets)
	sortByCreatedAtDesc(tweets)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"tweets": tweets,
	})
}

func getCurrentUserId(c *gin.Context) uint {

	// Benutzer aus dem Gin-Kontext abrufen
	userIdRaw, exists := c.Get("userId")
	if !exists {
		// Benutzer nicht gefunden
		return 0
	}

	userId, ok := userIdRaw.(uint)
	if !ok {
		// user id konnte nicht zu uint konvertiert werden
		return 0
	}

	return userId

}

// Sort descending by created at attribute
func sortByCreatedAtDesc(tweets []models.Tweet) {
	sort.Slice(tweets, func(i, j int) bool {
		return tweets[i].CreatedAt.After(tweets[j].CreatedAt)
	})
}
