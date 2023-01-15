package main

import (
	"blog_app/adapter/config"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/adapter/server"

	"github.com/gin-gonic/gin"
)

func main() {
	// load config
	config.Load()
	// setup logger

	// connect DB
	db := postgres.ConnectDB()

	engine := gin.Default()
	// DI
	registry := initialize(db)
	server.SetupRouter(engine, registry)
	// graceful shutdown
	port := ":" + config.Get().API.Port
	engine.Run(port)
}
