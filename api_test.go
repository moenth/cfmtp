package main

import (
    "log"
    "io/ioutil"
	"testing"

	"github.com/kataras/iris"
	"github.com/simplereach/timeutils"
)

func TestTradeGet(t *testing.T) {
	m := MgoDB{}
	m.Init()
    m.C("trades").DropCollection()
	defer m.Close()

	iris.Default = iris.New()
	iris.Config = iris.Default.Config
	iris.Config.Tester.ListeningAddr = "localhost:8080"
	iris.API("/api/v1/trades", TradeAPI{})
	e := iris.Tester(t)

	m.C("trades").Insert(NewTrade())
	t.Log("GET /api/v1/trades, expecting status 200, array of length 1")
	e.GET("/api/v1/trades").Expect().Status(iris.StatusOK).JSON().Array().Length().Equal(1)

    m.C("trades").DropCollection()
}

func TestTradePost(t *testing.T) {
	m := MgoDB{}
	m.Init()
    m.C("trades").DropCollection()
	defer m.Close()

    log.SetOutput(ioutil.Discard)
	iris.Default = iris.New()
	iris.Config = iris.Default.Config
	iris.Config.Tester.ListeningAddr = "localhost:8080"
	iris.API("/api/v1/trades", TradeAPI{})
	e := iris.Tester(t)

    trade := Trade{
        UserID: 123,
        CurrencyFrom: "EUR",
        CurrencyTo: "GBP",
        AmountSell: 1000,
        AmountBuy: 700,
        Rate: 0.7,
        TimePlaced: timeutils.Time{},
        OriginatingCountry: "IE",
    }

	t.Log("POST /api/v1/trades, expecting status 201")
	e.POST("/api/v1/trades").WithJSON(trade).Expect().Status(iris.StatusCreated)

    trade.AmountBuy = 5000
    t.Log("POST /api/v1/trades, expecting status 422")
	e.POST("/api/v1/trades").WithJSON(trade).Expect().Status(422)

    m.C("trades").DropCollection()
}

func TestTradeIndex(t *testing.T) {
    iris.Default = iris.New()
	iris.Config = iris.Default.Config
	iris.Config.Tester.ListeningAddr = "localhost:8080"
	iris.Get("/trades", TradeIndex)
	e := iris.Tester(t)

    e.GET("/trades").Expect().Status(iris.StatusOK).ContentType("text/html")
}
