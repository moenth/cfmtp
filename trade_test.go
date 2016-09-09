package main

import (
	"encoding/json"
	"testing"

	"github.com/simplereach/timeutils"
)

func TestTradeValidation(t *testing.T) {
	t.Log("Validating valid trade...")
	trade := Trade{
		UserID:             1,
		CurrencyFrom:       "EUR",
		CurrencyTo:         "GBP",
		AmountSell:         100,
		AmountBuy:          70,
		Rate:               0.7,
		TimePlaced:         timeutils.Time{},
		OriginatingCountry: "IE",
	}

	_, err := trade.Validate()
	if err != nil {
		t.Fatalf("Unexpected error validating trade: %v", err)
	}

	t.Log("Validating invalid trade...")
	trade = Trade{
		UserID:             0,
		CurrencyFrom:       "still",
		CurrencyTo:         "water",
		AmountSell:         -100,
		AmountBuy:          -70,
		Rate:               -0.7,
		TimePlaced:         timeutils.Time{},
		OriginatingCountry: "matt damon",
	}

	_, err = trade.Validate()
	if err == nil {
		t.Fatalf("Expected error, but got none")
	}
}

func TestTradeUnmarshal(t *testing.T) {
	t.Log("Unmarshaling trade from json...")

	var trade Trade
	bs := []byte(`{"userId": "134256", "currencyFrom": "EUR", "currencyTo": "GBP", "amountSell": 1000, "amountBuy": 747.10, "rate": 0.7471, "timePlaced" : "24-JAN-16 10:27:44", "originatingCountry" : "FR"}`)
	err := json.Unmarshal(bs, &trade)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
