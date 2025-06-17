package order_details

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	domainOrderDetails "github.com/iqbalnzls/watchcommerce/src/domain"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
)

type orderDetailsRepo struct {
	db *sql.DB
}

func NewOrderDetailsRepository(db *sql.DB) RepositoryIFace {
	if db == nil {
		panic("db connection is nil")
	}

	return &orderDetailsRepo{
		db: db,
	}
}

func (r *orderDetailsRepo) SaveBulkWithDBTrx(appCtx *appContext.AppContext, tx *sql.Tx, orderID int64, domains []domainOrderDetails.OrderDetails) (err error) {
	var (
		values    = make([]string, 0)
		valueArgs = make([]interface{}, 0)
		query     = `INSERT INTO commerce.order_details(order_id, product_id, quantity, price) VALUES %s`
	)

	for i, v := range domains {
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))
		valueArgs = append(valueArgs, orderID, v.ProductID, v.Quantity, v.Price)
	}

	_, err = tx.Exec(fmt.Sprintf(query, strings.Join(values, ",")), valueArgs...)
	if err != nil {
		err = errors.New(constant.ErrorDatabaseProblem)
		return
	}

	return
}

func (r *orderDetailsRepo) GetByOrderID(appCtx *appContext.AppContext, orderID int64) (domains []*domainOrderDetails.OrderDetails, err error) {
	var query = `SELECT id, order_id, product_id, quantity, price, created_at, updated_at FROM commerce.order_details WHERE order_id = $1`

	rows, err := r.db.Query(query, orderID)
	if err != nil {
		err = errors.New(constant.ErrorDatabaseProblem)
		return
	}

	for rows.Next() {
		var domain domainOrderDetails.OrderDetails
		if err = rows.Scan(&domain.ID, &domain.OrderID, &domain.ProductID, &domain.Quantity, &domain.Price, &domain.CreatedAt, &domain.UpdatedAt); err != nil {
			err = errors.New(constant.ErrorDatabaseProblem)
			return
		}

		domains = append(domains, &domain)
	}

	if len(domains) == 0 {
		err = errors.New(constant.ErrorDataNotFound)
		return
	}

	return
}
