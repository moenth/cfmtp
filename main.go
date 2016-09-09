package main

import (
	"github.com/kataras/iris"
)

func main() {
	
	// Get everything ready
	initDB()
	registerMiddleware()
	registerRoutes()

	// Serve app on :8080
	iris.Listen(":8080")
}

// initDB prepares the database for use.
func initDB() {
	m := MgoDB{}
	m.Init()
	defer m.Close()

	// Work with a fresh db for added convenience.
	m.DropDB(Database)
	m.Index(Database, []string{"_id"})
}

// registerMiddle registers middleware handlers with the router.
func registerMiddleware() {
	iris.Use(NewRateLimiter())
}

// registerRoutes registers endpoints with the router.
func registerRoutes() {

	// Serve a basic frontend.
	iris.Get("/trades", TradeIndex)

	// Register the trade api.
	iris.API("/api/v1/trades", TradeAPI{})
}
