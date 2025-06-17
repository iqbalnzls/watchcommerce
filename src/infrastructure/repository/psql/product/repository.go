package product

import (
	"database/sql"

	"github.com/iqbalnzls/watchcommerce/src/domain"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
)

type RepositoryIFace interface {
	Save(appCtx *appContext.AppContext, domain *domain.Product) (err error)
	UpdateByQuantityWithDBTrx(appCtx *appContext.AppContext, tx *sql.Tx, id, quantity int64) (err error)
	GetByID(appCtx *appContext.AppContext, id int64) (domain *domain.Product, err error)
	GetByBrandID(appCtx *appContext.AppContext, brandID int64) (domains []*domain.Product, err error)
}
