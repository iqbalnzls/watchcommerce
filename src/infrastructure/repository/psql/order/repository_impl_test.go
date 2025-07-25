package order_test

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	domainOrder "github.com/iqbalnzls/watchcommerce/src/domain"
	"github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/order"
	"github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/order_details"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
	"github.com/iqbalnzls/watchcommerce/src/shared/logger"
)

var appCtx = appContext.NewAppContext(&logger.Log{})

func TestNewOrderRepository(t *testing.T) {
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
			name: "init repository success",
			args: args{
				db: new(sql.DB),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = order.NewOrderRepository(tt.args.db)
				})
			} else {
				assert.NotPanics(t, func() {
					_ = order_details.NewOrderDetailsRepository(tt.args.db)
				})
			}
		})
	}
}

func Test_orderRepo_Get(t *testing.T) {
	type resp struct {
		err  error
		rows *sqlmock.Rows
	}
	type args struct {
		resp resp
	}

	var (
		tn    = time.Now()
		tests = []struct {
			name       string
			args       args
			wantDomain *domainOrder.Order
			wantErr    error
		}{
			{
				name: "error data not found",
				args: args{
					resp: resp{
						err: sql.ErrNoRows,
					},
				},
				wantErr: errors.New(constant.ErrorDataNotFound),
			},
			{
				name: "error database problem",
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
						rows: sqlmock.NewRows([]string{"id", "total", "created_at", "updated_at"}).AddRow(1, 12000, tn, tn),
					},
				},
				wantDomain: &domainOrder.Order{
					ID:        1,
					Total:     12000,
					CreatedAt: tn,
					UpdatedAt: tn,
				},
			},
		}
	)

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := order.NewOrderRepository(db)
			if tt.wantErr != nil {
				mock.ExpectQuery(`^SELECT (.+)FROM (.+)order`).WillReturnError(tt.args.resp.err)
			} else {
				mock.ExpectQuery(`^SELECT (.+)FROM (.+)order`).WithArgs(sqlmock.AnyArg()).WillReturnRows(tt.args.resp.rows)
			}
			domain, err := r.Get(appCtx, 1)
			assert.Equal(t, tt.wantDomain, domain)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_orderRepo_SaveWithDBTrx(t *testing.T) {
	type resp struct {
		err  error
		rows *sqlmock.Rows
	}
	type args struct {
		resp resp
	}

	var (
		domain = &domainOrder.Order{
			ID:        1,
			Total:     12,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		tests = []struct {
			name    string
			args    args
			wantId  int64
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
				wantId:  0,
			},
			{
				name: "success",
				args: args{
					resp: resp{
						rows: sqlmock.NewRows([]string{"id"}).AddRow(2),
					},
				},
				wantId: 2,
			},
		}
	)

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := order.NewOrderRepository(db)

			mock.ExpectBegin()
			if tt.wantErr != nil {
				mock.ExpectQuery(`^INSERT INTO (.+)order`).WillReturnError(tt.args.resp.err)
			} else {
				mock.ExpectQuery(`^INSERT INTO (.+)order`).WithArgs(sqlmock.AnyArg()).WillReturnRows(tt.args.resp.rows)
			}

			tx, err := db.Begin()
			assert.NoError(t, err)

			id, err := r.SaveWithDBTrx(appCtx, tx, domain)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantId, id)
		})
	}
}

func Test_orderRepo_BeginDBTrx(t *testing.T) {
	var (
		tests = []struct {
			name string
		}{
			{
				name: "success",
			},
		}
	)

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectBegin()

			r := order.NewOrderRepository(db)
			_, err := r.BeginDBTrx()
			assert.NoError(t, err)
		})
	}
}

func Test_orderRepo_CommitDBTrx(t *testing.T) {
	var (
		tests = []struct {
			name string
		}{
			{
				name: "success",
			},
		}
	)

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectCommit()

			tx, err := db.Begin()
			assert.NoError(t, err)

			r := order.NewOrderRepository(db)
			wantEr := r.CommitDBTrx(appCtx, tx)
			assert.NoError(t, wantEr)
		})
	}
}

func Test_orderRepo_RollbackDBTrx(t *testing.T) {
	var (
		tests = []struct {
			name string
		}{
			{
				name: "success",
			},
		}
	)

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectRollback()

			tx, err := db.Begin()
			assert.NoError(t, err)

			r := order.NewOrderRepository(db)
			err = r.RollbackDBTrx(appCtx, tx)
			assert.NoError(t, err)
		})
	}
}
