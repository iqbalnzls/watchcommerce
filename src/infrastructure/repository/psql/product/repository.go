package product

import (
	"database/sql"

	"github.com/iqbalnzls/watchcommerce/src/domain"
)

type ProductRepositoryIFace interface {
	Save(domain *domain.Product) (err error)
	UpdateByQuantityWithDBTrx(tx *sql.Tx, id, quantity int64) (err error)
	GetByID(id int64) (domain *domain.Product, err error)
	GetByBrandID(brandID int64) (domains []*domain.Product, err error)
}
