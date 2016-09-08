package trade

import (
	"github.com/moenth/cfmtp/db"
	"gopkg.in/mgo.v2"
)

// Repository provides access to trades.
type Repository struct {
	*db.MgoDB
}

// Store persists a trade in the database.
func (r Repository) Store(t Trade) (err error) {
	_, err = r.col().UpsertId(t.ID, t)
	return
}

// List returns a list of the last n trades.
func (r Repository) List(n int) (trades []Trade, err error) {
	err = r.col().Find(nil).Sort("-_id").Limit(n).All(&trades)
	return
}

// Count returns the total number of trades.
func (r Repository) Count() (n int, err error) {
	return r.col().Count()
}

func (r Repository) col() *mgo.Collection {
	return r.MgoDB.C("trades")
}
