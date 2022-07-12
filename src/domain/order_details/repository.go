package order_details

import "database/sql"

type OrderDetailsRepositoryIFace interface {
	SaveBulkWithDBTrx(tx *sql.Tx, orderID int64, domains []OrderDetails) (err error)
	GetByOrderID(orderID int64) (domains []*OrderDetails, err error)
}
