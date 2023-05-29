package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/controllers"
	"github.com/marcbudd/linkup-service/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User Routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// Tweet Routes
	r.POST("/tweet/post", middleware.RequireAuth, controllers.PostTweet)
	r.DELETE("/tweet/delete/:tweetId", middleware.RequireAuth, controllers.DeleteTweet)
	r.GET("/tweet/getByUser/:userId", controllers.GetTweetsByUserId)

	return r
}
