package brand_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/iqbalnzls/watchcommerce/src/dto"
	"github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/brand"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
	"github.com/iqbalnzls/watchcommerce/src/shared/logger"
	mocksPsql "github.com/iqbalnzls/watchcommerce/src/shared/mock/infrastructure/repository/psql"
	usecaseBrand "github.com/iqbalnzls/watchcommerce/src/usecase/brand"
)

var appCtx = appContext.NewAppContext(&logger.Log{})

func TestNewBrandService(t *testing.T) {
	type args struct {
		brandRepo brand.RepositoryIFace
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name:      "brand repository is nil",
			wantPanic: true,
		},
		{
			name: "init service success",
			args: args{
				brandRepo: &mocksPsql.BrandRepositoryIFaceMock{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = usecaseBrand.NewBrandService(tt.args.brandRepo)
				})
			} else {
				assert.NotPanics(t, func() {
					_ = usecaseBrand.NewBrandService(tt.args.brandRepo)
				})
			}
		})
	}
}

func Test_service_Save(t *testing.T) {
	type save struct {
		err error
	}
	type brandRepo struct {
		save save
	}
	type resp struct {
		brandRepo brandRepo
	}
	type args struct {
		resp resp
	}

	var (
		req = &dto.CreateBrandRequest{
			Name: "test",
		}
		tests = []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: "save brand error",
				args: args{
					resp: resp{
						brandRepo: brandRepo{
							save: save{
								err: errors.New(constant.ErrorDatabaseProblem),
							},
						},
					},
				},
				wantErr: true,
			},
			{
				name: "success",
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			brandRepo := &mocksPsql.BrandRepositoryIFaceMock{}
			brandRepo.On("Save", mock.Anything, mock.Anything).Return(tt.args.resp.brandRepo.save.err)

			s := usecaseBrand.NewBrandService(brandRepo)
			if err := s.Save(appCtx, req); (err != nil) != tt.wantErr {
				t.Errorf("SaveWithDBTrx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
