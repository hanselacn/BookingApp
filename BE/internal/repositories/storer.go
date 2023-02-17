package repositories

import "BookingApp/BE/pkg/mariadb"

type storer struct {
	dbTx mariadb.Transaction
}

// NewStore create ne instance database transaction
func NewStore(dbTx mariadb.Transaction) *storer {
	return &storer{dbTx: dbTx}
}
