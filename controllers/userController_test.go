package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/controllers"
	"github.com/stretchr/testify/assert"
)

// func mockDB() *gorm.DB {
// 	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

// 	db.AutoMigrate(&models.Comment{})
// 	db.AutoMigrate(&models.Follow{})
// 	db.AutoMigrate(&models.Message{})
// 	db.AutoMigrate(&models.Post{})
// 	db.AutoMigrate(&models.Token{})
// 	db.AutoMigrate(&models.User{})
// 	return db
// }

// func TestSignup_Success(t *testing.T) {

// // Define test data
// email := "john.doe@example.com"
// password := "strongPassword123!"
// username := "johndoe"
// birthDate := time.Now()
// name := "John Doe"
// bio := "Hello, I'm John Doe"

// userCreateDTO := models.UserCreateRequestDTO{
// 	Email:     email,
// 	Username:  username,
// 	Password:  password,
// 	BirthDate: birthDate,
// 	Name:      name,
// 	Bio:       bio,
// 	Image:     nil,
// }

// // Set up router
// gin.SetMode(gin.TestMode)
// router := gin.Default()

// // Set mock database
// initalizers.DB = mockDB()
// var user models.User
// result := initalizers.DB.Save(&user)
// if result.Error != nil {
// 	t.Error(result.Error)
// }

// // Perform a POST request to /api/users/signup with a sample JSON payload
// requestBody := userCreateDTO
// requestBodyBytes, _ := json.Marshal(requestBody)
// request, _ := http.NewRequest("POST", "/api/users/signup", bytes.NewBuffer(requestBodyBytes))
// request.Header.Set("Content-Type", "application/json")

// // Create a response recorder to capture the output
// responseRecorder := httptest.NewRecorder()

// // Call the Signup handler function
// router.POST("/api/users/signup", controllers.Signup)
// router.ServeHTTP(responseRecorder, request)

// // Check the response status code
// assert.Equal(t, http.StatusCreated, responseRecorder.Code)

// fmt.Println(responseRecorder.Body.String())

// // Check the response body
// var responseBody models.UserGetResponseDTO
// err := json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody)
// assert.Nil(t, err)
// assert.Equal(t, userCreateDTO.Username, responseBody.Username)
// assert.Equal(t, userCreateDTO.BirthDate, responseBody.BirthDate)
// assert.Equal(t, userCreateDTO.BirthDate, responseBody.Name)
// assert.Equal(t, userCreateDTO.Bio, responseBody.Bio)
// assert.Nil(t, responseBody.Image)
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
