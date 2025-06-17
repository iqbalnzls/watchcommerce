package order

import (
	"database/sql"

	"github.com/iqbalnzls/watchcommerce/src/domain"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
)

type RepositoryIFace interface {
	BeginDBTrx() (tx *sql.Tx, err error)
	CommitDBTrx(appCtx *appContext.AppContext, tx *sql.Tx) (err error)
	RollbackDBTrx(appCtx *appContext.AppContext, tx *sql.Tx) (err error)
	SaveWithDBTrx(appCtx *appContext.AppContext, tx *sql.Tx, domain *domain.Order) (id int64, err error)
	Get(appCtx *appContext.AppContext, id int64) (domain *domain.Order, err error)
}
