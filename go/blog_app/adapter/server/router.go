package server

import (
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

	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/signup", authHandler.Signup)
		authGroup.POST("/login", authHandler.Login)
	}

	postGroup := v1.Group("/posts")
	{
		postGroup.GET("", postHandler.GetPosts)
	}

	commentGroup := v1.Group("/comments")
	{
		commentGroup.POST("", commentHandler.CreateComment)
		commentGroup.PUT("", commentHandler.UpdateComment)
	}

	tagGroup := v1.Group("/tags")
	{
		tagGroup.GET("", tagHandler.GetTags)
	}
}
