package psql_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	domainProduct "github.com/iqbalnzls/watchcommerce/src/domain/product"
	"github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql"
	"github.com/iqbalnzls/watchcommerce/src/pkg/constant"
)

func TestNewProductRepository(t *testing.T) {
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
					_ = psql.NewProductRepository(tt.args.db)
				})
			} else {
				assert.NotPanics(t, func() {
					_ = psql.NewProductRepository(tt.args.db)
				})
			}
		})
	}
}

func Test_productRepo_GetByBrandID(t *testing.T) {
	type resp struct {
		err  error
		rows *sqlmock.Rows
	}
	type args struct {
		resp resp
	}

	var (
		tn      = time.Now()
		brandID = int64(2)
		tests   = []struct {
			name        string
			args        args
			wantDomains []*domainProduct.Product
			wantErr     error
			wantArgs    bool
		}{
			{
				name: "query error",
				args: args{
					resp: resp{
						err: errors.New("general error"),
					},
				},
				wantErr: errors.New(constant.ErrorDatabaseProblem),
			},
			{
				name: "scan error",
				args: args{
					resp: resp{
						rows: sqlmock.NewRows([]string{"id"}).AddRow("12"),
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
						rows: sqlmock.NewRows([]string{"id", "brand_id", "name", "price", "quantity", "created_at", "updated_at"}).
							AddRow(1, 12, "daytona", 1200, 12, tn, tn),
					},
				},
				wantArgs: true,
				wantDomains: []*domainProduct.Product{
					{
						ID:        1,
						BrandID:   12,
						Name:      "daytona",
						Price:     1200,
						Quantity:  12,
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
			r := psql.NewProductRepository(db)

			if tt.wantErr != nil && !tt.wantArgs {
				mock.ExpectQuery(`^SELECT (.+) FROM (.+)product`).WillReturnError(tt.args.resp.err)
			} else {
				mock.ExpectQuery(`^SELECT (.+) FROM (.+)product`).WithArgs(sqlmock.AnyArg()).WillReturnRows(tt.args.resp.rows)
			}

			gotDomains, err := r.GetByBrandID(brandID)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantDomains, gotDomains)
		})
	}
}

func Test_productRepo_GetByID(t *testing.T) {
	type resp struct {
		err  error
		rows *sqlmock.Rows
	}
	type args struct {
		resp resp
	}

	var (
		tn    = time.Now()
		id    = int64(123)
		tests = []struct {
			name       string
			args       args
			wantDomain *domainProduct.Product
			wantErr    error
		}{
			{
				name: "record not found",
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
						rows: sqlmock.NewRows([]string{"id", "brand_id", "name", "price", "quantity", "created_at", "updated_at"}).
							AddRow(12, 34, "g-shock", 1200, 12, tn, tn),
					},
				},
				wantDomain: &domainProduct.Product{
					ID:        12,
					BrandID:   34,
					Name:      "g-shock",
					Price:     1200,
					Quantity:  12,
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
			r := psql.NewProductRepository(db)

			if tt.wantErr != nil {
				mock.ExpectQuery(`^SELECT (.+) FROM (.+)product`).WillReturnError(tt.args.resp.err)
			} else {
				mock.ExpectQuery(`^SELECT (.+) FROM (.+)product`).WithArgs(sqlmock.AnyArg()).WillReturnRows(tt.args.resp.rows)
			}

			gotDomain, err := r.GetByID(id)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantDomain, gotDomain)
		})
	}
}

func Test_productRepo_Save(t *testing.T) {
	type resp struct {
		err    error
		result driver.Result
	}
	type args struct {
		resp resp
	}

	var (
		domain = &domainProduct.Product{
			ID:      12,
			BrandID: 2,
		}
		tests = []struct {
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
			r := psql.NewProductRepository(db)

			if tt.wantErr != nil {
				mock.ExpectExec(`^INSERT INTO (.+)product`).WillReturnError(tt.args.resp.err)
			} else {
				mock.ExpectExec(`^INSERT INTO (.+)product`).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(tt.args.resp.result)
			}

			err := r.Save(domain)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_productRepo_UpdateByQuantity(t *testing.T) {
	type resp struct {
		err    error
		result driver.Result
	}
	type args struct {
		resp resp
	}

	var (
		id       = int64(1)
		quantity = int64(12)
		tests    = []struct {
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
			r := psql.NewProductRepository(db)

			if tt.wantErr != nil {
				mock.ExpectExec(`^UPDATE (.+)product`).WillReturnError(tt.args.resp.err)
			} else {
				mock.ExpectExec(`^UPDATE (.+)product`).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(tt.args.resp.result)
			}

			err := r.UpdateByQuantity(id, quantity)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
