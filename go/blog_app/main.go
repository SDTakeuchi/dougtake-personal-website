package main

import (
	"blog_app/adapter/config"
	"blog_app/adapter/log"
	"blog_app/adapter/domain_impl/model/auth"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/adapter/server"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	// setup logger
	logger := log.NewLogger()

	// connect DB
	db := postgres.ConnectDB()

	jwSecretKey := config.Get().Token.SecretKey
	jwtIssuer, err := auth.NewJWTIssuer(jwSecretKey)
	if err != nil {
		logger.Fatalf("initializing JWTIssuer Failed: %v", err.Error())
		panic(err)
	}

	engine := gin.Default()
	// DI
	registry := initialize(db, jwtIssuer, logger)
	server.SetupRouter(engine, registry)
	// graceful shutdown
	port := ":" + config.Get().API.Port
	engine.Run(port)
}
