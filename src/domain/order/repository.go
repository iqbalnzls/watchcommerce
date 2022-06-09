package order

type OrderRepositoryIFace interface {
	Save(domain *Order) (id int64, err error)
	Get(id int64) (domain *Order, err error)
}
