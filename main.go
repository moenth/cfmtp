package main

import (
	"github.com/kataras/iris"
	"github.com/moenth/cfmtp/db"
	"github.com/moenth/cfmtp/trade"
	"github.com/moenth/cfmtp/middleware"
)

func main() {
	initDB()
	registerMiddleware()
	registerAPI()

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

// registerAPI registers api routes with the router.
func registerAPI() {
	iris.API("/api/v1/trades", trade.API{})
}
