package order

import (
	"database/sql"

	"github.com/iqbalnzls/watchcommerce/src/domain"
)

type OrderRepositoryIFace interface {
	BeginDBTrx() (tx *sql.Tx, err error)
	CommitDBTrx(tx *sql.Tx) (err error)
	RollbackDBTrx(tx *sql.Tx) (err error)
	SaveWithDBTrx(tx *sql.Tx, domain *domain.Order) (id int64, err error)
	Get(id int64) (domain *domain.Order, err error)
}
