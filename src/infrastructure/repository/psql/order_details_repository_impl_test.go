package psql_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	domainOrderDetails "github.com/iqbalnzls/watchcommerce/src/domain/order_details"
	"github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql"
	"github.com/iqbalnzls/watchcommerce/src/pkg/constant"
)

func TestNewOrderDetailsRepository(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name:      "db connection is nil",
			wantPanic: true,
		},
		{
			name: "success",
			args: args{
				db: new(sql.DB),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = psql.NewOrderDetailsRepository(tt.args.db)
				})
			} else {
				assert.NotPanics(t, func() {
					_ = psql.NewOrderDetailsRepository(tt.args.db)
				})
			}
		})
	}
}

func Test_orderDetailsRepo_GetByOrderID(t *testing.T) {
	type resp struct {
		err  error
		rows *sqlmock.Rows
	}
	type args struct {
		resp resp
	}

	var (
		tn      = time.Now()
		orderID = int64(12)
		tests   = []struct {
			name        string
			args        args
			wantDomains []*domainOrderDetails.OrderDetails
			wantErr     error
			wantArgs    bool
		}{
			{
				name: "error query",
				args: args{
					resp: resp{
						err: errors.New("general error"),
					},
				},
				wantErr: errors.New(constant.ErrorDatabaseProblem),
			},
			{
				name: "error scan",
				args: args{
					resp: resp{
						rows: sqlmock.NewRows([]string{"id", "brandID"}).AddRow("12", "11"),
					},
				},
				wantArgs: true,
				wantErr:  errors.New(constant.ErrorDatabaseProblem),
			},
			{
				name: "record not found",
				args: args{
					resp: resp{
						rows: sqlmock.NewRows([]string{}),
					},
				},
				wantArgs: true,
				wantErr:  errors.New(constant.ErrorDataNotFound),
			},
			{
				name: "success",
				args: args{
					resp: resp{
						rows: sqlmock.NewRows([]string{"id", "order_id", "product_id", "quantity", "price", "created_at", "updated_at"}).
							AddRow(1, 3, 12, 1, 100, tn, tn),
					},
				},
				wantArgs: true,
				wantDomains: []*domainOrderDetails.OrderDetails{
					{
						ID:        1,
						OrderID:   3,
						ProductID: 12,
						Quantity:  1,
						Price:     100,
						CreatedAt: tn,
						UpdatedAt: tn,
					},
				},
			},
		}
	)

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := psql.NewOrderDetailsRepository(db)

			if tt.wantErr != nil && !tt.wantArgs {
				mock.ExpectQuery(`^SELECT (.+)FROM (.+)order_details`).WillReturnError(tt.args.resp.err)
			} else {
				mock.ExpectQuery(`^SELECT (.+)FROM (.+)order_details`).WithArgs(sqlmock.AnyArg()).WillReturnRows(tt.args.resp.rows)
			}

			domain, err := r.GetByOrderID(orderID)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantDomains, domain)
		})
	}
}

func Test_orderDetailsRepo_SaveBulkWithDBTrx(t *testing.T) {

	type resp struct {
		err    error
		result driver.Result
	}
	type args struct {
		resp resp
	}
	var (
		domains = []domainOrderDetails.OrderDetails{
			{
				ID:      int64(1),
				OrderID: int64(2),
			},
		}
		orderID = int64(213)
		tests   = []struct {
			name    string
			args    args
			wantErr error
		}{
			{
				name: "error database",
				args: args{
					resp: resp{
						err: errors.New("error database"),
					},
				},
				wantErr: errors.New(constant.ErrorDatabaseProblem),
			},
			{
				name: "success",
				args: args{
					resp: resp{
						result: sqlmock.NewResult(1, 1),
					},
				},
			},
		}
	)

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := psql.NewOrderDetailsRepository(db)
			mock.ExpectBegin()
			if tt.wantErr != nil {
				mock.ExpectExec(`^INSERT INTO (.+)order_details`).WillReturnError(errors.New("general error"))
			} else {
				mock.ExpectExec(`^INSERT INTO (.+)order_details`).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(tt.args.resp.result)
			}

			tx, err := db.Begin()
			assert.NoError(t, err)

			err = r.SaveBulkWithDBTrx(tx, orderID, domains)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
