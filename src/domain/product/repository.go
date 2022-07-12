package product

import "database/sql"

type ProductRepositoryIFace interface {
	Save(domain *Product) (err error)
	UpdateByQuantityWithDBTrx(tx *sql.Tx, id, quantity int64) (err error)
	GetByID(id int64) (domain *Product, err error)
	GetByBrandID(brandID int64) (domains []*Product, err error)
}
