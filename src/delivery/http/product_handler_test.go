package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	inHttp "github.com/iqbalnzls/watchcommerce/src/delivery/http"
	"github.com/iqbalnzls/watchcommerce/src/dto"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
	mocksUsecaseProduct "github.com/iqbalnzls/watchcommerce/src/shared/mock/usecase/product"
	"github.com/iqbalnzls/watchcommerce/src/shared/validator"
	"github.com/iqbalnzls/watchcommerce/src/usecase/product"
)

func TestNewProductHandler(t *testing.T) {
	type args struct {
		productService product.ServiceIFace
		v              *validator.DataValidator
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "product service is nil",
			args: args{
				v: validator.NewValidator(),
			},
			wantPanic: true,
		},
		{
			name: "validator is nil",
			args: args{
				productService: &mocksUsecaseProduct.ProductServiceIFaceMock{},
			},
			wantPanic: true,
		},
		{
			name: "init product handler success",
			args: args{
				productService: new(mocksUsecaseProduct.ProductServiceIFaceMock),
				v:              validator.NewValidator(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = inHttp.NewProductHandler(tt.args.productService, tt.args.v)
				})
			} else {
				assert.NotPanics(t, func() {
					_ = inHttp.NewProductHandler(tt.args.productService, tt.args.v)
				})
			}
		})
	}
}

func Test_productHandler_Get(t *testing.T) {
	type get struct {
		data dto.GetProductResponse
		err  error
	}
	type productService struct {
		get get
	}
	type resp struct {
		productService productService
	}
	type req struct {
		method     string
		queryParam map[string]string
	}
	type args struct {
		req  req
		resp resp
	}

	var (
		tests = []struct {
			name string
			args args
		}{
			{
				name: "invalid method",
				args: args{
					req: req{
						method: http.MethodPost,
					},
				},
			},
			{
				name: "validate request failed",
				args: args{
					req: req{
						method: http.MethodGet,
						queryParam: map[string]string{
							"id": "",
						},
					},
				},
			},
			{
				name: "call service error",
				args: args{
					req: req{
						method: http.MethodGet,
						queryParam: map[string]string{
							"id": "1",
						},
					},
					resp: resp{
						productService: productService{
							get: get{
								err: errors.New(constant.ErrorDatabaseProblem),
							},
						},
					},
				},
			},
			{
				name: "get product success",
				args: args{
					req: req{
						method: http.MethodGet,
						queryParam: map[string]string{
							"id": "1",
						},
					},
					resp: resp{
						productService: productService{
							get: get{
								data: dto.GetProductResponse{
									ID:      1,
									Name:    "cx-100",
									BrandID: 1,
								},
							},
						},
					},
				},
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productService := new(mocksUsecaseProduct.ProductServiceIFaceMock)
			productService.On("Get", mock.Anything, mock.Anything).Return(tt.args.resp.productService.get.data, tt.args.resp.productService.get.err)

			req := httptest.NewRequest(tt.args.req.method, "/api/v1/product/get", nil)
			q := req.URL.Query()
			for k, v := range tt.args.req.queryParam {
				q.Add(k, v)
			}

			req.URL.RawQuery = q.Encode()

			rec := httptest.NewRecorder()

			ctx := context.WithValue(req.Context(), constant.AppContext, appCtx)

			h := inHttp.NewProductHandler(productService, validator.NewValidator())
			h.Get(rec, req.WithContext(ctx))
		})
	}
}

func Test_productHandler_GetByBrandID(t *testing.T) {
	type getByBrandID struct {
		data []dto.GetProductResponse
		err  error
	}
	type productService struct {
		getByBrandID getByBrandID
	}
	type resp struct {
		productService productService
	}
	type req struct {
		method     string
		queryParam map[string]string
	}
	type args struct {
		req  req
		resp resp
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "invalid http method",
			args: args{
				req: req{
					method: http.MethodPost,
				},
			},
		},
		{
			name: "validate request failed",
			args: args{
				req: req{
					method: http.MethodGet,
					queryParam: map[string]string{
						"id": "",
					},
				},
			},
		},
		{
			name: "call service error",
			args: args{
				req: req{
					method: http.MethodGet,
					queryParam: map[string]string{
						"id": "1",
					},
				},
				resp: resp{
					productService: productService{
						getByBrandID: getByBrandID{
							err: errors.New(constant.ErrorDatabaseProblem),
						},
					},
				},
			},
		},
		{
			name: "call service error",
			args: args{
				req: req{
					method: http.MethodGet,
					queryParam: map[string]string{
						"id": "1",
					},
				},
				resp: resp{
					productService: productService{
						getByBrandID: getByBrandID{
							data: []dto.GetProductResponse{
								{
									ID:       1,
									Name:     "cx-900",
									BrandID:  2,
									Price:    1000,
									Quantity: 12,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productService := new(mocksUsecaseProduct.ProductServiceIFaceMock)
			productService.On("GetByBrandID", mock.Anything, mock.Anything).Return(tt.args.resp.productService.getByBrandID.data, tt.args.resp.productService.getByBrandID.err)

			req := httptest.NewRequest(tt.args.req.method, "/api/v1/product/get", nil)
			q := req.URL.Query()
			for k, v := range tt.args.req.queryParam {
				q.Add(k, v)
			}

			req.URL.RawQuery = q.Encode()

			rec := httptest.NewRecorder()

			ctx := context.WithValue(req.Context(), constant.AppContext, appCtx)

			h := inHttp.NewProductHandler(productService, validator.NewValidator())
			h.GetByBrandID(rec, req.WithContext(ctx))
		})
	}
}

func Test_productHandler_Save(t *testing.T) {
	type save struct {
		err error
	}
	type productService struct {
		save save
	}
	type resp struct {
		productService productService
	}
	type req struct {
		method  string
		payload interface{}
	}
	type args struct {
		req  req
		resp resp
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "invalid method",
			args: args{
				req: req{
					method: http.MethodGet,
				},
			},
		},
		{
			name: "decode payload error",
			args: args{
				req: req{
					method:  http.MethodPost,
					payload: "test",
				},
			},
		},
		{
			name: "validate payload failed",
			args: args{
				req: req{
					method: http.MethodPost,
					payload: dto.CreateProductRequest{
						BrandID:  01,
						Price:    1000,
						Quantity: 12,
					},
				},
			},
		},
		{
			name: "call service error",
			args: args{
				req: req{
					method: http.MethodPost,
					payload: dto.CreateProductRequest{
						BrandID:  2,
						Name:     "ls-200",
						Price:    9000,
						Quantity: 2,
					},
				},
				resp: resp{
					productService: productService{
						save: save{
							err: errors.New(constant.ErrorDatabaseProblem),
						},
					},
				},
			},
		},
		{
			name: "create product success",
			args: args{
				req: req{
					method: http.MethodPost,
					payload: dto.CreateProductRequest{
						BrandID:  3,
						Name:     "ls-300",
						Price:    11000,
						Quantity: 5,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productService := new(mocksUsecaseProduct.ProductServiceIFaceMock)
			productService.On("Save", mock.Anything, mock.Anything).Return(tt.args.resp.productService.save.err)

			b, _ := json.Marshal(tt.args.req.payload)

			req := httptest.NewRequest(tt.args.req.method, "/api/v1/product/get", bytes.NewReader(b))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			ctx := context.WithValue(req.Context(), constant.AppContext, appCtx)

			h := inHttp.NewProductHandler(productService, validator.NewValidator())
			h.Save(rec, req.WithContext(ctx))
		})
	}
}
