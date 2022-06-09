package product

type ProductRepositoryIFace interface {
	Save(domain *Product) (err error)
	UpdateByQuantity(id, quantity int64) (err error)
	GetByID(id int64) (domain *Product, err error)
	GetByBrandID(brandID int64) (domains []*Product, err error)
}
