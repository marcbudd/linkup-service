package router

import (
	"embed"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/marcbudd/linkup-service/controllers"
	"github.com/marcbudd/linkup-service/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var swagger embed.FS

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Set CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "https://link-up-rho.vercel.app"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// Set trusted proxies
	r.SetTrustedProxies([]string{os.Getenv("PROXY_HOST")})

	// Swagger
	r.StaticFS("/swaggerio", http.FS(swagger))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Redirect from "/swagger" to "/swagger/index.html"
	r.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// API Routes
	api := r.Group("/api")

	// User Routes
	api.POST("/users/signup", controllers.Signup)
	api.POST("/users/login", controllers.Login)
	api.GET("/users/validate", middleware.RequireAuth, controllers.Validate)
	api.PATCH("/users/confirm/:token", controllers.ConfirmEmail)
	api.PATCH("/users/updatePassword", middleware.RequireAuth, controllers.UpdatePassword)
	api.GET("/users/:userID", middleware.RequireAuth, controllers.GetUserByID)
	api.GET("/users/current", middleware.RequireAuth, controllers.GetCurrentUser)
	api.GET("/users", controllers.GetUsers)
	api.PATCH("/users", middleware.RequireAuth, controllers.UpdateUser)
	api.PATCH("/users/forgotPassword", controllers.UpdatePasswordForgotten)
	api.DELETE("users/delete", middleware.RequireAuth, controllers.DeleteUser)

	// Post Routes
	api.POST("/posts", middleware.RequireAuth, controllers.CreatePost)
	api.DELETE("/posts/:postID", middleware.RequireAuth, controllers.DeletePost)
	api.GET("/posts/:postID", middleware.RequireAuth, controllers.GetPostByID)
	api.GET("/posts/user/:userID", middleware.RequireAuth, controllers.GetPostsByUserID)
	api.GET("/posts/feed", middleware.RequireAuth, controllers.GetPostsForCurrentUser)
	api.GET("/posts", middleware.RequireAuth, controllers.GetPosts)
	api.GET("/posts/gpt", controllers.CreatePostGPT)

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
	api.POST("/comments", middleware.RequireAuth, controllers.CreateComment)
	api.DELETE("/comments/:commentID", middleware.RequireAuth, controllers.DeleteComment)
	api.GET("/comments/posts/:postID", controllers.GetCommentsByPostID)

	// Message Routes
	api.POST("/messages", middleware.RequireAuth, controllers.CreateMessage)
	api.GET("/messages/:chatPartnerID", middleware.RequireAuth, controllers.GetMessagesByChat)
	api.GET("/messages", middleware.RequireAuth, controllers.GetChatsByUserID)

	return r
}
