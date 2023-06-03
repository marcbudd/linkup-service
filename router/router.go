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
	api.POST("/post", middleware.RequireAuth, controllers.CreatePost)
	api.DELETE("/post/delete/:postId", middleware.RequireAuth, controllers.DeletePost)
	api.GET("/post/byUserId/:userId", controllers.GetPostsByUserId)

	// Like Routes
	api.POST("like/:postId", middleware.RequireAuth, controllers.CreateLike)
	api.DELETE("like/:postId", middleware.RequireAuth, controllers.DeleteLike) // By Post Id
	api.GET("like/byPostId/:postId", controllers.GetLikesByPostId)

	// Follow Routes
	api.POST("follow/:userId", middleware.RequireAuth, controllers.CreateFollow)
	// api.DELETE("follow/:userId")

	// Comment Routes
	api.POST("comment/:postId", middleware.RequireAuth, controllers.CreateComment)
	api.DELETE("comment/:commentId", middleware.RequireAuth, controllers.DeleteComment)
	api.GET("comment/byPostId/:postId", controllers.GetCommentsByPostId)

	// Test mail route
	api.GET("/testMail", controllers.SendMail)

	return r
}
