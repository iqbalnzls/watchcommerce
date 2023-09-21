package order_details

import (
	"database/sql"

	"github.com/iqbalnzls/watchcommerce/src/domain"
)

type OrderDetailsRepositoryIFace interface {
	SaveBulkWithDBTrx(tx *sql.Tx, orderID int64, domains []domain.OrderDetails) (err error)
	GetByOrderID(orderID int64) (domains []*domain.OrderDetails, err error)
}
