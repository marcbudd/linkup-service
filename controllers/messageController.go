package controllers

import (
	"github.com/gin-gonic/gin"
)

func CreateMessage(c *gin.Context) {

	// 	// Get content from body
	// 	var messageCreateRequestDTO models.MessageCreateRequestDTO

	// 	if c.Bind(&messageCreateRequestDTO) != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Failed to read body",
	// 		})
	// 	}

	// 	// Get user id of logged in user
	// 	userId := getCurrentUserId(c)
	// 	if userId == 0 {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Invalid user id",
	// 		})
	// 	}

	// 	// Create message
	// 	message := models.Message{
	// 		SenderID: userId,
	// 		ReceiverID: messageCreateRequestDTO.ReceiverID,
	// 		Text: messageCreateRequestDTO.Text,
	// 	}

	// 	result := initalizers.DB.Create(&message)

	// 	if result.Error != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Failed to create message",
	// 		})
	// 	}

	// 	// Respond
	// 	c.JSON(http.StatusCreated, gin.H{
	// 		message: message,
	// 	})

	// }

	// func GetMessages(c *gin.Context) {

	// 	// Get user id of param
	// 	userId := c.Param("userId")

	// 	// Get messages
	// 	var messagesReceived []models.Message
	// 	result := initalizers.DB.Where("receiver_id = ?", userId).Find(&messages)

	// 	if result.Error != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Failed to get messages",
	// 		})
	// 	}

	// 	var messagesSent []models.Message
	// 	result := initalizers.DB.Where("sender_id = ?", userId).Find(&messages)

	// 	if result.Error != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Failed to get messages",
	// 		})
	// 	}

	// 	var combinedMessages []models.Message
	// 	combinedMessages = append(messagesReceived, messagesSent)
	// 	combinedMessagesSorted = append(combinedMessages, func(i, j int) bool {
	// 		return combinedMessages[i].CreatedAt.Before(combinedMessages[j].CreatedAt)
	// 	})

	// 	// Respond
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"messages": combinedMessagesSorted,
	// 	})
}
