package main

import (
	"log"
)

// ProcessTrade processes an incoming trade.
func ProcessTrade(t Trade) {
	m := MgoDB{}
	m.Init()
	defer m.Close()

	// Store the trade in the database
	_, err := m.C("trades").UpsertId(t.ID, t)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("Processed trade: %s\n", t.ID)
}
