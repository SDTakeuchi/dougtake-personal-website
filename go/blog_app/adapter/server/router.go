package server

import (
	"blog_app/adapter/registry"

	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine, registry registry.Registry) {
	var (
		postHandler    = registry.PostHandler
		commentHandler = registry.CommentHandler
	)

	postGroup := engine.Group("/posts")
	{
		postGroup.GET("", postHandler.GetPosts)
	}

	commentGroup := engine.Group("/comments")
	{
		commentGroup.POST("", commentHandler.CreateComment)
	}
}

//func InitRouting(e *echo.Echo, tagProvider tagProvider) {
//	g := e.Group("/api")
//	g.POST("/presets/", tagProvider.Post())
//	g.GET("/presets/", tagProvider.Get())
//	g.GET("/presets/:id", tagProvider.FindByID())
//	g.PUT("/presets/:id", tagProvider.Put())
//	g.DELETE("/presets/:id", tagProvider.Delete())
//}
