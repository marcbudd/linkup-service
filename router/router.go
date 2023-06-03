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
	api.GET("/post/:postID", controllers.GetPostByID)
	api.GET("/post/byUserID/:userID", controllers.GetPostsByUserID)
	api.GET("/post/byCurrentUser", middleware.RequireAuth, controllers.GetPostsByCurrentUser)
	api.GET("/post/forCurrentUser", middleware.RequireAuth, controllers.GetPostsForCurrentUser)
	api.GET("/post", controllers.GetPosts)

	// Like Routes
	api.POST("like/:postID", middleware.RequireAuth, controllers.CreateLike)
	api.DELETE("like/:postID", middleware.RequireAuth, controllers.DeleteLike) // By Post Id
	api.GET("like/byPostId/:postID", controllers.GetLikesByPostId)

	// Follow Routes
	api.POST("follow/:followedUserID", middleware.RequireAuth, controllers.CreateFollow)
	api.DELETE("follow/:followedUserID", middleware.RequireAuth, controllers.DeleteFollow)
	api.GET("follow/followingsOfUserID/:userID", controllers.GetFollowingsOfUserID)
	api.GET("follow/followersOfUserID/:userID", controllers.GetFollowersOfUserID)

	// Comment Routes
	api.POST("comment/:postID", middleware.RequireAuth, controllers.CreateComment)
	api.DELETE("comment/:commentID", middleware.RequireAuth, controllers.DeleteComment)
	api.GET("comment/byPostId/:postID", controllers.GetCommentsByPostId)

	// Message Routes
	api.POST("message/:receiverID", middleware.RequireAuth, controllers.CreateMessage)
	api.GET("message/:receiverID", middleware.RequireAuth, controllers.GetMessagesByChat)

	return r
}
