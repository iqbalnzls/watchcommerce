package order_details

import (
	"database/sql"

	"github.com/iqbalnzls/watchcommerce/src/domain"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
)

type RepositoryIFace interface {
	SaveBulkWithDBTrx(appCtx *appContext.AppContext, tx *sql.Tx, orderID int64, domains []domain.OrderDetails) (err error)
	GetByOrderID(appCtx *appContext.AppContext, orderID int64) (domains []*domain.OrderDetails, err error)
}
