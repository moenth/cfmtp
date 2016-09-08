package trade

import (
	"log"

	"github.com/kataras/iris"
	"github.com/moenth/cfmtp/db"
)

// API exposes endpoints for market interactions.
type API struct {
	*iris.Context
}

// Get lists a number of trades.
// GET /trades
func (api API) Get() {
	db := db.MgoDB{}
	db.Init()
	defer db.Close()

    trades := Repository{&db}
	recent, err := trades.List(20)
	if err != nil {
		api.JSON(500, err.Error())
		return
	}

	api.JSON(200, recent)
}

// Post consumes an incoming trade.
// POST /trades
func (api API) Post() {
	trade := NewTrade()

	// Read the trade and verify the input is well-formed.
	err := api.ReadJSON(&trade)
	if err != nil {
		api.JSON(400, err.Error())
		return
	}

	// Perform sanity checks on the incoming trade.
	// Return 422 unprocessable entity on validation failure.
	_, err = trade.Validate()
	if err != nil {
		api.JSON(422, err.Error())
		return
	}

	// Accept trade for processing.
	go ProcessTrade(trade)
	log.Printf("Received trade: %v", trade)
	api.SetStatusCode(201)
}
