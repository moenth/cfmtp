package main

import (
	"github.com/kataras/iris"
	"github.com/moenth/cfmtp/db"
	"github.com/moenth/cfmtp/middleware"
	"github.com/moenth/cfmtp/trade"
)

func main() {
	iris.Config.IsDevelopment = true
	// Get everything ready
	initDB()
	registerMiddleware()
	registerRoutes()

	// Serve app on :8080
	iris.Listen(":8080")
}

// initDB prepares the database for use.
func initDB() {
	m := db.MgoDB{}
	m.Init()
	defer m.Close()

	// Work with a fresh db for convenience sake.
	m.DropDB(db.Database)
	m.Index(db.Database, []string{"_id"})
}

// registerMiddle registers middleware handlers with the router.
func registerMiddleware() {
	iris.Use(middleware.NewRateLimiter())
}

// registerRoutes registers endpoints with the router.
func registerRoutes() {

	// Serve a basic frontend.
	iris.Get("/trades", trade.Index)

	// Register the trade api.
	iris.API("/api/v1/trades", trade.API{})
}
