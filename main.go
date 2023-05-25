package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/controllers"
	"github.com/marcbudd/linkup-service/initalizers"
	"github.com/marcbudd/linkup-service/middleware"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDb()
	initalizers.SyncDatabase()
}

func main() {
	r := gin.Default()

	//user routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	//tweet routes
	r.POST("/tweet/post", middleware.RequireAuth, controllers.PostTweet)
	r.DELETE("/tweet/delete/:tweetId", middleware.RequireAuth, controllers.DeleteTweet)
	r.GET("/tweet/getByUser/:userId", controllers.GetTweetsByUserId)

	r.Run()
}
