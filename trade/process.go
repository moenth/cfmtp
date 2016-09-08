package trade

import (
	"log"

	"github.com/moenth/cfmtp/db"
)

// ProcessTrade processes an incoming trade.
func ProcessTrade(t Trade) {
	db := db.MgoDB{}
	db.Init()
	defer db.Close()

	// Store the trade in the database
	trades := Repository{&db}
	err := trades.Store(t)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("Processed trade: %s\n", t.ID)
}
