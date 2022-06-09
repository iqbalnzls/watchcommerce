package order_details

type OrderDetailsRepositoryIFace interface {
	SaveBulk(orderID int64, domains []OrderDetails) (err error)
	GetByOrderID(orderID int64) (domains []*OrderDetails, err error)
}
