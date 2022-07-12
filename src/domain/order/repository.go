package order

import "database/sql"

type OrderRepositoryIFace interface {
	BeginDBTrx() (tx *sql.Tx, err error)
	CommitDBTrx(tx *sql.Tx) (err error)
	RollbackDBTrx(tx *sql.Tx) (err error)
	SaveWithDBTrx(tx *sql.Tx, domain *Order) (id int64, err error)
	Get(id int64) (domain *Order, err error)
}
