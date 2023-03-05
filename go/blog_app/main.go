package main

import (
	"blog_app/adapter/config"
	"blog_app/adapter/domain_impl/model/auth"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/adapter/server"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	// setup logger

	// connect DB
	db := postgres.ConnectDB()

	jwSecretKey := config.Get().Auth.SecretKey
	jwtIssuer, err := auth.NewJWTIssuer(jwSecretKey)
	if err != nil {
		log.Fatalf("initializing JWTIssuer Failed: %v", err.Error())
		panic(err)
	}

	engine := gin.Default()
	// DI
	registry := initialize(db, jwtIssuer)
	server.SetupRouter(engine, registry)
	// graceful shutdown
	port := ":" + config.Get().API.Port
	engine.Run(port)
}
