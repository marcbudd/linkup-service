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
	api.POST("/user/signup", controllers.Signup)
	api.POST("/user/login", controllers.Login)
	api.GET("/user/validate", middleware.RequireAuth, controllers.Validate)
	api.PATCH("/user/confirmEmail/:token", middleware.RequireAuth, controllers.ConfirmEmail)
	api.PATCH("/user/changePassword", middleware.RequireAuth, controllers.UpdatePassword)
	api.GET("/user/:userID", controllers.GetUserByID)
	api.GET("/user", controllers.GetUsers)
	api.PATCH("/user", middleware.RequireAuth, controllers.UpdateUser)
	api.PATCH("/user/forgotPassword", controllers.UpdatePasswordForgotten)

	// Post Routes
	api.POST("/post", middleware.RequireAuth, controllers.CreatePost)
	api.DELETE("/post/delete/:postID", middleware.RequireAuth, controllers.DeletePost)
	api.GET("/post/byUserID/:userID", controllers.GetPostsByUserID)
	api.GET("/post/byCurrentUser", middleware.RequireAuth, controllers.GetPostsByCurrentUser)

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

	return r
}
