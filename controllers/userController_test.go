package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/controllers"
	"github.com/stretchr/testify/assert"
)

// func TestSignup_Success(t *testing.T) {
// 	// Setup
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.POST("/signup", controllers.Signup)

// 	// Create a valid userCreateDTO
// 	userCreateDTO := models.UserCreateRequestDTO{
// 		Email:    "test@example.com",
// 		Username: "testuser",
// 		Password: "password",
// 	}

// 	// TODO: Create mock database or mock functions

// 	// Perform the request
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/signup", nil)
// 	req.PostForm = userCreateDTO.ToValues()
// 	router.ServeHTTP(w, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusCreated, w.Code)
// 	assert.Equal(t, "mocked-token", w.Header().Get("Set-Cookie"))
// }

func TestSignup_FailedToReadBody(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/signup", controllers.Signup)

	// Perform the request without a valid userCreateDTO
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signup", nil)
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "failed to read body")
}

// func TestSignup_FailedToCreateUser(t *testing.T) {
// 	// Setup
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.POST("/signup", controllers.Signup)

// 	// TODO: Mock database or functions

// 	// Perform the request with a valid userCreateDTO
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/signup", nil)
// 	req.PostForm = userCreateDTO.ToValues()
// 	router.ServeHTTP(w, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusInternalServerError, w.Code)
// 	assert.Contains(t, w.Body.String(), "failed to create user")
// }

// func TestSignup_FailedToGenerateToken(t *testing.T) {
// 	// Setup
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.POST("/signup", controllers.Signup)

// 	// TODO: Mock database or functions

// 	// Perform the request with a valid userCreateDTO
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/signup", nil)
// 	req.PostForm = userCreateDTO.ToValues()
// 	router.ServeHTTP(w, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// 	assert.Contains(t, w.Body.String(), "Failed to create token")
// }
