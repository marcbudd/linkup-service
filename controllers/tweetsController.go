package controllers

import (
	"net/http"

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
	userId := getUserID(c)
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

func getUserID(c *gin.Context) uint {

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
