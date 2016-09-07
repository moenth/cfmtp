package trade

import (
    "github.com/moenth/cfmtp/db"
)

const (
    collection = "trades"
)

// Repository provides access to trades.
type Repository struct {
    *db.MgoDB
}

// Store persists a trade in the database.
func (r Repository) Store(t Trade) (err error) {
    _, err = r.MgoDB.C(collection).UpsertId(t.ID, t)
    return
}
