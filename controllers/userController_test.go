package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/controllers"
	"github.com/stretchr/testify/assert"
)

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
