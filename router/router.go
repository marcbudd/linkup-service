package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/controllers"
	"github.com/marcbudd/linkup-service/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	// User Routes
	api.POST("/signup", controllers.Signup)
	api.POST("/login", controllers.Login)
	api.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// Post Routes
	api.POST("/post/post", middleware.RequireAuth, controllers.CreatePost)
	api.DELETE("/post/delete/:postId", middleware.RequireAuth, controllers.DeletePost)
	api.GET("/post/byUserId/:userId", controllers.GetPostsByUserId)

	return r
}
