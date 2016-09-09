package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestProcessTrade(t *testing.T) {
	t.Log("Processing trade...")
	m := MgoDB{}
	m.Init()
	m.DropDB(Database)
	defer m.Close()

	// Surpress log output
	log.SetOutput(ioutil.Discard)

	trade := NewTrade()
	ProcessTrade(trade)
	count, _ := m.C("trades").Count()

	if count != 1 {
		t.Fatalf("Expected count of 1, got %d", count)
	}
}
