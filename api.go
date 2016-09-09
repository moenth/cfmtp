package main

import (
	"log"

	"github.com/kataras/iris"
)

// TradeAPI exposes endpoints for market interactions.
type TradeAPI struct {
	*iris.Context
}

// Get lists a number of trades.
// GET /trades
func (api TradeAPI) Get() {
	m := MgoDB{}
	m.Init()
	defer m.Close()

	var ts []Trade
	err := m.C("trades").Find(nil).Limit(20).All(&ts)
	if err != nil {
		api.JSON(500, err.Error())
		return
	}

	api.JSON(200, ts)
}

// Post consumes an incoming trade.
// POST /trades
func (api TradeAPI) Post() {
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

// TradeIndex renders the trades index page.
func TradeIndex(c *iris.Context) {
	m := MgoDB{}
	m.Init()
	defer m.Close()

	var ts []Trade
	m.C("trades").Find(nil).Limit(20).All(&ts)
	tc, _ := m.C("trades").Count()

	c.Render("trades.html", struct {
		Trades []Trade
		Total  int
	}{ts, tc})
}
