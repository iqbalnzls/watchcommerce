package brand

import (
	"database/sql"
	"errors"

	domainBrand "github.com/iqbalnzls/watchcommerce/src/domain"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
)

type brandRepo struct {
	db *sql.DB
}

func NewRepositoryBrand(db *sql.DB) RepositoryIFace {
	if db == nil {
		panic("db connection is nil")
	}

	return &brandRepo{
		db: db,
	}
}

func (r *brandRepo) Save(appCtx *appContext.AppContext, domain *domainBrand.Brand) (err error) {
	var query = `INSERT INTO commerce.brand(name) VALUES($1)`

	_, err = r.db.Exec(query, domain.Name)
	if err != nil {
		appCtx.Logger.Error(err.Error())
		err = errors.New(constant.ErrorDatabaseProblem)
		return
	}

	return
}
