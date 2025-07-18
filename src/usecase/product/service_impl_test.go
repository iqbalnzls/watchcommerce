package product_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	domainProduct "github.com/iqbalnzls/watchcommerce/src/domain"
	"github.com/iqbalnzls/watchcommerce/src/dto"
	"github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/product"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
	"github.com/iqbalnzls/watchcommerce/src/shared/logger"
	mocksPsql "github.com/iqbalnzls/watchcommerce/src/shared/mock/infrastructure/repository/psql"
	usecaseProduct "github.com/iqbalnzls/watchcommerce/src/usecase/product"
)

var appCtx = appContext.NewAppContext(&logger.Log{})

func TestNewProductService(t *testing.T) {
	type args struct {
		productRepo product.RepositoryIFace
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name:      "product repository is nil",
			wantPanic: true,
		},
		{
			name: "init service success",
			args: args{
				productRepo: new(mocksPsql.ProductRepositoryIFaceMock),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = usecaseProduct.NewProductService(tt.args.productRepo)
				})
			} else {
				assert.NotPanics(t, func() {
					_ = usecaseProduct.NewProductService(tt.args.productRepo)
				})
			}
		})
	}
}

func Test_productService_Get(t *testing.T) {
	type getByID struct {
		domain *domainProduct.Product
		err    error
	}
	type productRepo struct {
		getByID getByID
	}
	type resp struct {
		productRepo productRepo
	}
	type args struct {
		resp
	}

	var (
		req = &dto.GetProductRequest{
			ProductID: 1,
		}
		tests = []struct {
			name     string
			args     args
			wantResp dto.GetProductResponse
			wantErr  bool
		}{
			{
				name: "call Get func error",
				args: args{
					resp: resp{
						productRepo: productRepo{
							getByID: getByID{
								err: errors.New(constant.ErrorDataNotFound),
							},
						},
					},
				},
				wantErr: true,
			},
			{
				name: "success",
				args: args{
					resp: resp{
						productRepo: productRepo{
							getByID: getByID{
								domain: &domainProduct.Product{
									ID:   1,
									Name: "daytona",
								},
							},
						},
					},
				},
				wantResp: dto.GetProductResponse{
					ID:   1,
					Name: "daytona",
				},
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := new(mocksPsql.ProductRepositoryIFaceMock)
			productRepo.On("GetByID", mock.Anything, mock.Anything).Return(tt.args.productRepo.getByID.domain, tt.args.productRepo.getByID.err)

			s := usecaseProduct.NewProductService(productRepo)
			gotResp, err := s.Get(appCtx, req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Get() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_productService_GetByBrandID(t *testing.T) {
	type getByBrandID struct {
		domains []*domainProduct.Product
		err     error
	}
	type productRepo struct {
		getByBrandID getByBrandID
	}
	type resp struct {
		productRepo productRepo
	}
	type args struct {
		resp resp
	}

	var (
		req = &dto.GetProductByBrandIDRequest{
			BrandID: 2,
		}
		tests = []struct {
			name     string
			args     args
			wantResp []dto.GetProductResponse
			wantErr  bool
		}{
			{
				name: "cal GetByBrandID func error",
				args: args{
					resp: resp{
						productRepo: productRepo{
							getByBrandID: getByBrandID{
								err: errors.New(constant.ErrorDatabaseProblem),
							},
						},
					},
				},
				wantErr: true,
			},
			{
				name: "success",
				args: args{
					resp: resp{
						productRepo: productRepo{
							getByBrandID: getByBrandID{
								domains: []*domainProduct.Product{
									{
										ID:      1,
										BrandID: 1,
										Name:    "cx-100",
										Brand: domainProduct.Brand{
											Name: "rolex",
										},
									},
									{
										ID:      2,
										BrandID: 1,
										Name:    "cx-200",
										Brand: domainProduct.Brand{
											Name: "rolex",
										},
									},
								},
							},
						},
					},
				},
				wantResp: []dto.GetProductResponse{
					{
						ID:        1,
						Name:      "cx-100",
						BrandName: "rolex",
					},
					{
						ID:        2,
						Name:      "cx-200",
						BrandName: "rolex",
					},
				},
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := new(mocksPsql.ProductRepositoryIFaceMock)
			productRepo.On("GetByBrandID", mock.Anything, mock.Anything).Return(tt.args.resp.productRepo.getByBrandID.domains, tt.args.resp.productRepo.getByBrandID.err)

			s := usecaseProduct.NewProductService(productRepo)
			gotResp, err := s.GetByBrandID(appCtx, req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByBrandID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetByBrandID() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_productService_Save(t *testing.T) {
	type save struct {
		err error
	}
	type productRepo struct {
		save save
	}
	type resp struct {
		productRepo productRepo
	}
	type args struct {
		resp resp
	}

	var (
		req = &dto.CreateProductRequest{
			BrandID:  1,
			Name:     "rolex",
			Price:    1000,
			Quantity: 12,
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
						productRepo: productRepo{
							save: save{
								err: errors.New(constant.ErrorDatabaseProblem),
							},
						},
					},
				},
				wantErr: true,
			},
			{
				name: "save brand success",
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := new(mocksPsql.ProductRepositoryIFaceMock)
			productRepo.On("Save", mock.Anything, mock.Anything).Return(tt.args.resp.productRepo.save.err)

			s := usecaseProduct.NewProductService(productRepo)
			if err := s.Save(appCtx, req); (err != nil) != tt.wantErr {
				t.Errorf("SaveWithDBTrx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
