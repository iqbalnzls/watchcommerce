package psql

import (
	"database/sql"
	"errors"

	domainBrand "github.com/iqbalnzls/watchcommerce/src/domain/brand"
	"github.com/iqbalnzls/watchcommerce/src/pkg/constant"
	"github.com/iqbalnzls/watchcommerce/src/pkg/utils"
)

type brandRepo struct {
	db *sql.DB
}

func NewRepositoryBrand(db *sql.DB) domainBrand.BrandRepositoryIFace {
	if db == nil {
		panic("db connection is nil")
	}

	return &brandRepo{
		db: db,
	}
}

func (r *brandRepo) Save(domain *domainBrand.Brand) (err error) {
	var query = `INSERT INTO commerce.brand(name) VALUES($1)`

	_, err = r.db.Exec(query, domain.Name)
	if err != nil {
		utils.Error(err)
		err = errors.New(constant.ErrorDatabaseProblem)
		return
	}

	return
}
