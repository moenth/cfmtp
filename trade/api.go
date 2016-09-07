package trade

import (
	"log"

	"github.com/kataras/iris"
)

// API exposes endpoints for market interactions.
type API struct {
	*iris.Context
}

// Post consumes an incoming trade.
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
	api.JSON(201, nil)
}
