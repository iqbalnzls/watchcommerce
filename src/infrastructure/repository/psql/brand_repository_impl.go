package brand

import (
	"database/sql"

	"github.com/iqbalnzls/watchcommerce/src/domain/brand"
)

type brandRepo struct {
	db *sql.DB
}

func NewRepositoryBrand(db *sql.DB) brand.RepositoryBrandIFace {
	if db == nil {
		panic("db connection is nil")
	}

	return &brandRepo{
		db: db,
	}
}

func (b brandRepo) Save(domain *brand.Brand) (err error) {
	//TODO implement me
	panic("implement me")
}
