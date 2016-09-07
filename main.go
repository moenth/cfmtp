package main

import (
	"github.com/kataras/iris"
	"github.com/moenth/cfmtp/trade"
)

func main() {
	iris.API("/api/v1/trades", trade.API{})
	iris.Listen(":8080")
}
