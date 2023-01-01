package main

import "blog_app/adapter/persistance/database/postgres"

func main() {
	// load config

	// setup logger

	// connect DB
	db := postgres.ConnectDB()

	// DI
	e := initialize()

	// graceful shutdown
}
