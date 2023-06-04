package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/services"
)

func CreateMessage(c *gin.Context) {

	// Read body
	var messageCreateRequestDTO models.MessageCreateRequestDTO
	if c.Bind(&messageCreateRequestDTO) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
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

	// Create message
	err := services.CreateMessage(userID.(uint), messageCreateRequestDTO)

	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, gin.H{})
}

func GetMessagesByChat(c *gin.Context) {

	// Get user id of logged in user
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Get post id from url
	chatPartnerID := c.Param("chatPartnerID")
	if chatPartnerID == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Get messages
	messages, err := services.GetMessagesByChat(userID.(uint), chatPartnerID)

	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, messages)
}
