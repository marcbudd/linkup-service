package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcbudd/twitter-clone-backend/controllers"
	"github.com/marcbudd/twitter-clone-backend/initalizers"
	"github.com/marcbudd/twitter-clone-backend/middleware"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDb()
	initalizers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
