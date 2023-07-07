package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/services"
)

// CreateMessage creates a new message.
// @Summary Create a new message
// @Description Creates a new message sent by the logged-in user
// @Tags Messages
// @Accept json
// @Produce json
// @Param messageCreateRequestDTO body models.MessageCreateRequestDTO true "Message data"
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/messages [post]
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
	message, err := services.CreateMessage(userID.(uint), messageCreateRequestDTO)

	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, message)
}

// GetMessagesByChat retrieves messages between the logged-in user and a chat partner.
// @Summary Get messages by chat
// @Description Retrieves messages between the logged-in user and a chat partner
// @Tags Messages
// @Param chatPartnerID path string true "ID of the chat partner"
// @Produce json
// @Success 200 {array} models.MessagesOfChatGetResponseDTO
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/messages/{chatPartnerID} [get]
func GetMessagesByChat(c *gin.Context) {

	// Get user id of logged in user
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Get id from chat partner from url
	chatPartnerID := c.Param("chatPartnerID")
	if chatPartnerID == "" {
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

// GetChatsByUserID retrieves chats associated with the logged-in user.
// @Summary Get chats by user ID
// @Description Retrieves chats associated with the logged-in user
// @Tags Messages
// @Produce json
// @Success 200 {array} models.ChatOfUserGetResponseDTO
// @Failure 401
// @Failure 500
// @Router /api/messages [get]
func GetChatsByUserID(c *gin.Context) {
	// Get user id of logged in user
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Get chats
	chats, err := services.GetChatsByUserID(userID.(uint))
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, chats)

}
