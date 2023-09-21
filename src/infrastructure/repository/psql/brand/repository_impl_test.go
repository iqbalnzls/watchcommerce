package brand_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	domainBrand "github.com/iqbalnzls/watchcommerce/src/domain"
	"github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/brand"
	"github.com/iqbalnzls/watchcommerce/src/pkg/constant"
)

func TestNewRepositoryBrand(t *testing.T) {
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
					_ = brand.NewRepositoryBrand(tt.args.db)
				})
			} else {
				assert.NotPanics(t, func() {
					_ = brand.NewRepositoryBrand(tt.args.db)
				})
			}
		})
	}
}

func Test_brandRepo_Save(t *testing.T) {
	type resp struct {
		err    error
		result driver.Result
	}
	type args struct {
		resp resp
	}

	var (
		domain = &domainBrand.Brand{
			ID:        0,
			Name:      "",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
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
			r := brand.NewRepositoryBrand(db)
			if tt.wantErr != nil {
				mock.ExpectExec(`^INSERT INTO (.+)brand`).WillReturnError(tt.args.resp.err)
			} else {
				mock.ExpectExec(`^INSERT INTO (.+)brand`).WithArgs(sqlmock.AnyArg()).WillReturnResult(tt.args.resp.result)
			}
			err = r.Save(domain)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
