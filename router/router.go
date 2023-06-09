package router

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/controllers"
	"github.com/marcbudd/linkup-service/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var swagger embed.FS

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Swagger
	r.StaticFS("/swaggerio", http.FS(swagger))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API Routes
	api := r.Group("/api")

	// User Routes
	api.POST("/users/signup", controllers.Signup)
	api.POST("/users/login", controllers.Login)
	api.GET("/users/validate", middleware.RequireAuth, controllers.Validate)
	api.PATCH("/users/confirm/:token", controllers.ConfirmEmail)
	api.PATCH("/users/updatePassword", middleware.RequireAuth, controllers.UpdatePassword)
	api.GET("/users/:userID", controllers.GetUserByID)
	api.GET("/users", controllers.GetUsers)
	api.PATCH("/users", middleware.RequireAuth, controllers.UpdateUser)
	api.PATCH("/users/forgotPassword", controllers.UpdatePasswordForgotten)

	// Post Routes
	api.POST("/posts", middleware.RequireAuth, controllers.CreatePost)
	api.DELETE("/posts/:postID", middleware.RequireAuth, controllers.DeletePost)
	api.GET("/posts/:postID", controllers.GetPostByID)
	api.GET("/posts/user/:userID", controllers.GetPostsByUserID)
	api.GET("/posts/feed", middleware.RequireAuth, controllers.GetPostsForCurrentUser)
	api.GET("/posts", controllers.GetPosts)

	// Like Routes
	api.POST("/likes/:postID", middleware.RequireAuth, controllers.CreateLike)
	api.DELETE("/likes/:postID", middleware.RequireAuth, controllers.DeleteLike) // By Post Id
	api.GET("/likes/:postID", controllers.GetLikesByPostId)

	// Follow Routes
	api.POST("/follows/:followedUserID", middleware.RequireAuth, controllers.CreateFollow)
	api.DELETE("/follows/:followedUserID", middleware.RequireAuth, controllers.DeleteFollow)
	api.GET("/follows/:userID/followings", controllers.GetFollowingsByUserID)
	api.GET("/follows/:userID/followers", controllers.GetFollowersByUserID)

	// Comment Routes
	api.POST("/comments/:postID", middleware.RequireAuth, controllers.CreateComment)
	api.DELETE("/comments/:commentID", middleware.RequireAuth, controllers.DeleteComment)
	api.GET("/comments/posts/:postID", controllers.GetCommentsByPostID)

	// Message Routes
	api.POST("/messages", middleware.RequireAuth, controllers.CreateMessage)
	api.GET("/messages/:chatPartnerID", middleware.RequireAuth, controllers.GetMessagesByChat)
	api.GET("/messages", middleware.RequireAuth, controllers.GetChatsByUserID)

	return r
}
