package server

import (
	"blog_app/adapter/handler/middleware"
	"blog_app/adapter/registry"

	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine, registry registry.Registry) {
	var (
		authHandler    = registry.AuthHandler
		postHandler    = registry.PostHandler
		commentHandler = registry.CommentHandler
		tagHandler     = registry.TagHandler

		apiGroup = engine.Group("/api")
		v1       = apiGroup.Group("/v1")
	)

	// authedRoutes := v1.Group("/").Use(middleware.CheckAuth(registry.TokenIssuer))

	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/signup", authHandler.Signup)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/renew_token", authHandler.RenewToken)
	}

	postGroup := v1.Group("/posts")
	{
		postGroup.GET("", postHandler.GetPosts)
		postGroup.POST("", postHandler.CreatePost)
	}

	commentGroup := v1.Group("/comments").Use(middleware.CheckAuth(registry.TokenIssuer))
	{
		commentGroup.POST("", commentHandler.CreateComment)
		commentGroup.PUT("", commentHandler.UpdateComment)
		commentGroup.DELETE("", commentHandler.DeleteComment)
	}

	tagGroup := v1.Group("/tags")
	{
		tagGroup.GET("", tagHandler.GetTags)
	}
}
