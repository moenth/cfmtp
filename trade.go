package main

import (
	"errors"
	"time"

	"github.com/simplereach/timeutils"
	"gopkg.in/mgo.v2/bson"
)

// Trade contains information about a single market interaction.
type Trade struct {
	ID                 string         `json:"-" bson:"_id"`
	UserID             int            `json:"userId,string" bson:"user_id"`
	CurrencyFrom       string         `json:"currencyFrom" bson:"currency_from"`
	CurrencyTo         string         `json:"currencyTo" bson:"currency_to"`
	AmountSell         float32        `json:"amountSell" bson:"amount_sell"`
	AmountBuy          float32        `json:"amountBuy" bson:"amount_buy"`
	Rate               float32        `json:"rate" bson:"rate"`
	TimePlaced         timeutils.Time `json:"timePlaced" bson:"time_placed"`
	OriginatingCountry string         `json:"originatingCountry" bson:"originating_country"`
}

// Validate checks the trade for validity.
func (t Trade) Validate() (valid bool, err error) {

	if t.UserID <= 0 {
		err = errors.New("User ID is not valid")
	}

	// We only trade in positive amounts
	if t.AmountBuy <= 0 || t.AmountSell <= 0 {
		err = errors.New("Amount buy and amount sell must be positive")
		return
	}

	// Same with the rate
	if t.Rate <= 0 {
		err = errors.New("Rate must be positive")
		return
	}

	// Amounts and rate should match up
	if t.AmountSell*t.Rate != t.AmountBuy {
		err = errors.New("Trade amounts do not match trade rate")
		return
	}

	// Trades from the future are prohibited
	if t.TimePlaced.Time.After(time.Now()) {
		err = errors.New("Trade cannot take place in the future")
		return
	}

	valid = true
	return
}

// NewTrade creates a new trade.
func NewTrade() Trade {
	return Trade{
		ID: bson.NewObjectId().Hex(),
	}
}
