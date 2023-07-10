package controllers

import (
	"net/http"
	"strconv"

	"github.com/marcbudd/linkup-service/models"
	"github.com/marcbudd/linkup-service/services"
	"github.com/marcbudd/linkup-service/utils"

	"github.com/gin-gonic/gin"
)

// Signup creates a new user account
// @Summary User Signup
// @Description Create a new user account
// @Tags Users
// @Accept json
// @Produce json
// @Param userCreateDTO body models.UserCreateRequestDTO true "User information"
// @Success 201 {object} models.UserGetResponseDTO
// @Failure 400
// @Failure 500
// @Router /api/users/signup [post]
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
	c.SetCookie("Authorization", tokenString, 3600, "/", "", true, true)
	c.JSON(http.StatusCreated, user)

}

// Login authenticates and logs in a user
// @Summary User Login
// @Description Authenticate and login a user
// @Tags Users
// @Accept json
// @Param userLoginRequestDTO body models.UserLoginRequestDTO true "User credentials"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/users/login [post]
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
	c.SetCookie("Authorization", tokenString, 3600, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{})

}

// Validate validates the user token and checks if the user is authorized
// @Summary Validate User Token
// @Description Validate the user token and check if the user is authorized
// @Tags Users
// @Security ApiKeyAuth
// @Success 200
// @Failure 401
// @Router /api/users/validate [get]
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

// ConfirmEmail godoc
// @Summary Confirm Email
// @Description Confirm the user's email address using the provided token
// @Tags Users
// @Accept json
// @Produce json
// @Param token path string true "Confirmation token"
// @Success 200
// @Failure 401
// @Failure 404
// @Router /api/users/confirm/{token} [patch]
func ConfirmEmail(c *gin.Context) {
	// Get token from url
	token := c.Param("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Confirm email
	err := services.ConfirmEmail(token)
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})

}

// UpdatePassword updates the password of the logged-in user
// @Summary Update Password
// @Description Update the password of the logged-in user
// @Tags Users
// @Accept json
// @Param password body models.UserUpdatePasswortRequestDTO true "New password"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /api/users/updatePassword [patch]
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

// GetUserByID returns a user by their ID
// @Summary Get User by ID
// @Description Get a user by their ID
// @Tags Users
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {object} models.UserDetailGetResponseDTO
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /api/users/{userID} [get]
func GetUserByID(c *gin.Context) {

	// Get parameter from url
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

	// Get user
	user, err := services.GetUserByID(userID, currentUserID.(uint))
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, user)
}

// GetCurrentUser returns the logged-in user
// @Summary Get Current User
// @Description Get the logged-in user
// @Tags Users
// @Produce json
// @Success 200 {object} models.UserDetailGetResponseDTO
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /api/users/current [get]
func GetCurrentUser(c *gin.Context) {

	// Get user id of logged in user
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Get user
	user, err := services.GetUserByID(strconv.Itoa(int(userID.(uint))), userID.(uint))
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, user)
}

// GetUsers returns a list of users based on query parameters
// @Summary Get Users
// @Description Get a list of users based on query parameters
// @Tags Users
// @Accept json
// @Produce json
// @Param query query string false "Search query"
// @Param limit query int false "Limit the number of results per page"
// @Param page query int false "Page number"
// @Success 200 {array} models.UserGetResponseDTO
// @Failure 500
// @Router /api/users [get]
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

// UpdateUser updates the information of a user
// @Summary Update user
// @Description Update the user's information
// @Tags Users
// @Accept json
// @Param userID path int true "User ID"
// @Param userUpdate body models.UserUpdateRequestDTO true "User update data"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 409
// @Failure 500
// @Router /api/users [patch]
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

// UpdatePasswordForgotten resets the forgotten password
// @Summary Reset forgotten password
// @Description Reset the forgotten password
// @Tags Users
// @Accept json
// @Param updatePasswordForgotten body models.UserUpdatePasswordForgottenRequestDTO true "User update password forgotten data"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/users/forgotPassword [patch]
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

// DeleteUser deletes the logged-in user
// @Summary Delete user
// @Description Delete the logged-in user
// @Tags Users
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /api/users [delete]
func DeleteUser(c *gin.Context) {
	// Get user id of logged in user
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		return
	}

	// Delete user
	err := services.DeleteUser(userID.(uint))
	if err != nil {
		c.JSON(err.HTTPStatusCode(), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}
