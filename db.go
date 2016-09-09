package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

// Connection details
const (
	Host     = "mongodb://db:27017"
	Database = "cfmtp"
)

// Main mgo connections
var (
	mainSession *mgo.Session
	mainDb      *mgo.Database
)

// MgoDB wraps the main mgo session.
type MgoDB struct {
	Session *mgo.Session
	DB      *mgo.Database
	Col     *mgo.Collection
}

// Initialize the main session on init.
func init() {
	if mainSession == nil {
		var err error
		mainSession, err = mgo.Dial(Host)

		if err != nil {
			log.Fatal(err)
		}

		mainSession.SetMode(mgo.Monotonic, true)
		mainDb = mainSession.DB(Database)
	}
}

// Init requests a new socket connection from the connection pool.
func (m *MgoDB) Init() *mgo.Session {
	m.Session = mainSession.Copy()
	m.DB = m.Session.DB(Database)

	return m.Session
}

// C returns a value representing the named collection.
func (m *MgoDB) C(name string) *mgo.Collection {
	m.Col = m.DB.C(name)
	return m.Col
}

// DropDB drops the specified database.
func (m *MgoDB) DropDB(name string) {
	err := m.Session.DB(name).DropDatabase()
	if err != nil {
		log.Fatal(err)
	}
}

// Index ensures keys on the given collection are indexed.
func (m *MgoDB) Index(collection string, keys []string) {
	index := mgo.Index{
		Key:        keys,
		Unique:     true,
		DropDups:   true,
		Background: true,
	}

	err := m.C(collection).EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}
}

// Close closes the connection and returns it to the pool.
func (m *MgoDB) Close() {
	defer m.Session.Close()
}
