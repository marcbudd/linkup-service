package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var tokenExpirationHours = 24 * 30 // token expires after 30 days

func Signup(c *gin.Context) {

	// Get email, username and password of body
	var userSignupDTO models.UserSignupRequestDTO

	if c.Bind(&userSignupDTO) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(userSignupDTO.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash Password",
		})

		return
	}

	// Create user
	user := models.User{
		Username:     userSignupDTO.Username,
		Email:        userSignupDTO.Email,
		PasswordHash: string(hash),
	}
	result := initalizers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, gin.H{})

}

func Login(c *gin.Context) {

	// Get email and password of body
	var userLoginRequestDTO models.UserLoginRequestDTO

	if c.Bind(&userLoginRequestDTO) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Look up requested user
	var user models.User
	initalizers.DB.First(&user, "email= ?", userLoginRequestDTO.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Compare sent in passqord with saved user password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userLoginRequestDTO.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * time.Duration(tokenExpirationHours)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	// Respond
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*tokenExpirationHours, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})

}

func Validate(c *gin.Context) {
	userId, _ := c.Get("userId")

	c.JSON(http.StatusOK, gin.H{
		"message": userId,
	})
}
