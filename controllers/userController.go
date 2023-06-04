package controllers

import (
	"net/http"
	"strconv"

	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/services"
	"github.com/marcbudd/linkup-service/utils"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	// Read body
	var userCreateDTO models.UserCreateRequestDTO

	if c.Bind(&userCreateDTO) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	// Create user
	user, serviceErr := services.CreateUser(userCreateDTO)
	if serviceErr != nil {
		c.JSON(serviceErr.HTTPStatusCode(), gin.H{
			"error": serviceErr.Error(),
		})
		return
	}

	// Generate a jwt token
	tokenString, err := utils.GenerateJWTToken(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	// Respond
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*30*24, "", "", false, true)
	c.JSON(http.StatusCreated, gin.H{})

}

func Login(c *gin.Context) {

	// Read body
	var userLoginRequestDTO models.UserLoginRequestDTO

	if c.Bind(&userLoginRequestDTO) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Look up requested user
	user, serviceErr := services.LoginUser(userLoginRequestDTO)
	if serviceErr != nil {
		c.JSON(serviceErr.HTTPStatusCode(), gin.H{
			"error": serviceErr.Error(),
		})
	}

	// Generate a jwt token
	tokenString, err := utils.GenerateJWTToken(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	// Respond
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*30*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})

}

func Validate(c *gin.Context) {
	// Get user id of logged in user
	_, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "validated",
	})
}

func ConfirmEmail(c *gin.Context) {
	// Get user id of logged in user
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Confirm email
	err := services.ConfirmEmail(userID.(uint), c.Param("token"))
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})

}

func UpdatePassword(c *gin.Context) {

	// Read body
	var updatePasswordDTO models.UserUpdatePasswortRequestDTO

	if c.Bind(&updatePasswordDTO) != nil {
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

	// Update password
	err := services.UpdatePassword(userID.(uint), updatePasswordDTO)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})

}

func GetUserByID(c *gin.Context) {

	// Get parameter from url
	userID := c.Param("userID")

	// Get user
	user, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {

	// Get query paramters
	query := c.Query("query")
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 0
	}
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		page = 0
	}

	// Get users
	users, serviceErr := services.GetUsers(query, int(limit), int(page))
	if serviceErr != nil {
		c.JSON(serviceErr.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, users)

}

func UpdateUser(c *gin.Context) {
	// Read body
	var updateUpdateDTO models.UserUpdateRequestDTO

	if c.Bind(&updateUpdateDTO) != nil {
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

	// Update user
	err := services.UpdateUser(userID.(uint), updateUpdateDTO)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

func UpdatePasswordForgotten(c *gin.Context) {
	// Read body
	var updatePasswordForgottenDTO models.UserUpdatePasswordForgottenRequestDTO

	if c.Bind(&updatePasswordForgottenDTO) != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Reset password
	err := services.UpdatePasswordForgotten(updatePasswordForgottenDTO)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})

}
