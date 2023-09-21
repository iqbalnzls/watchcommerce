package order

import (
	"database/sql"
	"errors"

	domainOrder "github.com/iqbalnzls/watchcommerce/src/domain"
	"github.com/iqbalnzls/watchcommerce/src/pkg/constant"
	"github.com/iqbalnzls/watchcommerce/src/pkg/utils"
)

type orderRepo struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepositoryIFace {
	if db == nil {
		panic("db connection is nil")
	}

	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) SaveWithDBTrx(tx *sql.Tx, domain *domainOrder.Order) (id int64, err error) {
	var query = `INSERT INTO commerce."order"(total) VALUES($1) RETURNING id`

	if err = tx.QueryRow(query, domain.Total).Scan(&id); err != nil {
		utils.Error(err)
		err = errors.New(constant.ErrorDatabaseProblem)
		return
	}
	return
}

func (r *orderRepo) Get(id int64) (domain *domainOrder.Order, err error) {
	var (
		query = `SELECT id, total, created_at, updated_at FROM commerce."order" WHERE id = $1`
		order domainOrder.Order
	)

	if err = r.db.QueryRow(query, id).Scan(&order.ID, &order.Total, &order.CreatedAt, &order.UpdatedAt); err != nil {
		utils.Error(err)
		if err == sql.ErrNoRows {
			err = errors.New(constant.ErrorDataNotFound)
			return
		}
		err = errors.New(constant.ErrorDatabaseProblem)
		return
	}

	domain = &order

	return
}

func (r *orderRepo) BeginDBTrx() (tx *sql.Tx, err error) {
	tx, err = r.db.Begin()
	if err != nil {
		err = errors.New(constant.ErrorDatabaseProblem)
		return
	}
	return
}

func (r *orderRepo) CommitDBTrx(tx *sql.Tx) (err error) {
	if err = tx.Commit(); err != nil {
		err = errors.New(constant.ErrorDatabaseProblem)
		return
	}
	return
}

func (r *orderRepo) RollbackDBTrx(tx *sql.Tx) (err error) {
	if err = tx.Rollback(); err != nil {
		err = errors.New(constant.ErrorDatabaseProblem)
		return
	}
	return
}
