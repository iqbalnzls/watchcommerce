package order_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	domainOrder "github.com/iqbalnzls/watchcommerce/src/domain/order"
	domainOrderDetails "github.com/iqbalnzls/watchcommerce/src/domain/order_details"
	domainProduct "github.com/iqbalnzls/watchcommerce/src/domain/product"
	"github.com/iqbalnzls/watchcommerce/src/dto"
	"github.com/iqbalnzls/watchcommerce/src/pkg/constant"
	mocksPsql "github.com/iqbalnzls/watchcommerce/src/pkg/mock/infrastructure/repository/psql"
	usecaseOrder "github.com/iqbalnzls/watchcommerce/src/usecase/order"
)

func TestNewOrderService(t *testing.T) {
	type args struct {
		productRepo      domainProduct.ProductRepositoryIFace
		orderRepo        domainOrder.OrderRepositoryIFace
		orderDetailsRepo domainOrderDetails.OrderDetailsRepositoryIFace
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "product repository is nil",
			args: args{
				orderRepo:        new(mocksPsql.OrderRepositoryIFaceMock),
				orderDetailsRepo: new(mocksPsql.OrderDetailsRepositoryIFaceMock),
			},
			wantPanic: true,
		},
		{
			name: "order repository is nil",
			args: args{
				productRepo:      new(mocksPsql.ProductRepositoryIFaceMock),
				orderDetailsRepo: new(mocksPsql.OrderDetailsRepositoryIFaceMock),
			},
			wantPanic: true,
		},
		{
			name: "order details repository is nil",
			args: args{
				productRepo: new(mocksPsql.ProductRepositoryIFaceMock),
				orderRepo:   new(mocksPsql.OrderRepositoryIFaceMock),
			},
			wantPanic: true,
		},
		{
			name: "init service success",
			args: args{
				productRepo:      new(mocksPsql.ProductRepositoryIFaceMock),
				orderRepo:        new(mocksPsql.OrderRepositoryIFaceMock),
				orderDetailsRepo: new(mocksPsql.OrderDetailsRepositoryIFaceMock),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = usecaseOrder.NewOrderService(tt.args.productRepo, tt.args.orderRepo, tt.args.orderDetailsRepo)
				})
			} else {
				assert.NotPanics(t, func() {
					_ = usecaseOrder.NewOrderService(tt.args.productRepo, tt.args.orderRepo, tt.args.orderDetailsRepo)
				})
			}
		})
	}
}

func Test_orderService_Get(t *testing.T) {
	type getByOrderID struct {
		domains []*domainOrderDetails.OrderDetails
		err     error
	}
	type orderDetailsRepo struct {
		getByOrderID getByOrderID
	}
	type get struct {
		domain *domainOrder.Order
		err    error
	}
	type orderRepo struct {
		get get
	}
	type resp struct {
		orderRepo        orderRepo
		orderDetailsRepo orderDetailsRepo
	}
	type args struct {
		resp resp
	}

	var (
		req = &dto.GetOrderRequest{
			OrderID: 1,
		}
		tests = []struct {
			name     string
			args     args
			wantResp dto.GetOrderResponse
			wantErr  bool
		}{
			{
				name: "call Get func error",
				args: args{
					resp: resp{
						orderRepo: orderRepo{
							get: get{
								err: errors.New(constant.ErrorDatabaseProblem),
							},
						},
					},
				},
				wantErr: true,
			},
			{
				name: "call GetByOrderID func error",
				args: args{
					resp: resp{
						orderRepo: orderRepo{
							get: get{
								domain: &domainOrder.Order{
									ID:    1,
									Total: 1000,
								},
							},
						},
						orderDetailsRepo: orderDetailsRepo{
							getByOrderID: getByOrderID{
								err: errors.New(constant.ErrorDatabaseProblem),
							},
						},
					},
				},
				wantErr: true,
			},
			{
				name: "get order success",
				args: args{
					resp: resp{
						orderRepo: orderRepo{
							get: get{
								domain: &domainOrder.Order{
									ID:    1,
									Total: 1000,
								},
							},
						},
						orderDetailsRepo: orderDetailsRepo{
							getByOrderID: getByOrderID{
								domains: []*domainOrderDetails.OrderDetails{
									{
										ID:        1,
										OrderID:   12,
										ProductID: 1,
									},
								},
							},
						},
					},
				},
				wantResp: dto.GetOrderResponse{
					ID:    1,
					Total: 1000,
					Details: []dto.OrderDetailsResponse{
						{
							ID:        1,
							OrderID:   12,
							ProductID: 1,
						},
					},
				},
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := new(mocksPsql.ProductRepositoryIFaceMock)

			orderRepo := new(mocksPsql.OrderRepositoryIFaceMock)
			orderRepo.On("Get", mock.Anything).Return(tt.args.resp.orderRepo.get.domain, tt.args.resp.orderRepo.get.err)

			orderDetailsRepo := new(mocksPsql.OrderDetailsRepositoryIFaceMock)
			orderDetailsRepo.On("GetByOrderID", mock.Anything).Return(tt.args.resp.orderDetailsRepo.getByOrderID.domains, tt.args.resp.orderDetailsRepo.getByOrderID.err)

			s := usecaseOrder.NewOrderService(productRepo, orderRepo, orderDetailsRepo)
			gotResp, err := s.Get(req)
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

func Test_orderService_Save(t *testing.T) {
	type getByID struct {
		domain *domainProduct.Product
		err    error
	}
	type updateByQuantity struct {
		err error
	}
	type productRepo struct {
		getByID          getByID
		updateByQuantity updateByQuantity
	}
	type save struct {
		id  int64
		err error
	}
	type orderRepo struct {
		save save
	}
	type saveBulk struct {
		err error
	}
	type orderDetailsRepo struct {
		saveBulk saveBulk
	}
	type resp struct {
		productRepo      productRepo
		orderRepo        orderRepo
		orderDetailsRepo orderDetailsRepo
	}
	type args struct {
		resp resp
		req  *dto.CreateOrderRequest
	}

	var (
		tests = []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: "error database when look up product",
				args: args{
					resp: resp{
						productRepo: productRepo{
							getByID: getByID{
								err: errors.New(constant.ErrorDatabaseProblem),
							},
						},
					},
					req: &dto.CreateOrderRequest{
						OrderDetails: []dto.OrderDetailsRequest{
							{
								ProductID: 1,
								Quantity:  2,
							},
						},
					},
				},
				wantErr: true,
			},
			{
				name: "product quantity less than request quantity",
				args: args{
					resp: resp{
						productRepo: productRepo{
							getByID: getByID{
								domain: &domainProduct.Product{
									ID:       1,
									Quantity: 1,
								},
							},
						},
					},
					req: &dto.CreateOrderRequest{
						OrderDetails: []dto.OrderDetailsRequest{
							{
								ProductID: 1,
								Quantity:  2,
							},
						},
					},
				},
				wantErr: true,
			},
			{
				name: "save order error",
				args: args{
					resp: resp{
						productRepo: productRepo{
							getByID: getByID{
								domain: &domainProduct.Product{
									ID:       1,
									Quantity: 10,
								},
							},
						},
						orderRepo: orderRepo{
							save: save{
								err: errors.New(constant.ErrorDatabaseProblem),
							},
						},
					},
					req: &dto.CreateOrderRequest{
						OrderDetails: []dto.OrderDetailsRequest{
							{
								ProductID: 1,
								Quantity:  2,
							},
						},
					},
				},
				wantErr: true,
			},
			{
				name: "save bulk error",
				args: args{
					resp: resp{
						productRepo: productRepo{
							getByID: getByID{
								domain: &domainProduct.Product{
									ID:       1,
									Quantity: 10,
								},
							},
						},
						orderRepo: orderRepo{
							save: save{
								id: 12,
							},
						},
						orderDetailsRepo: orderDetailsRepo{
							saveBulk: saveBulk{
								err: errors.New(constant.ErrorDatabaseProblem),
							},
						},
					},
					req: &dto.CreateOrderRequest{
						OrderDetails: []dto.OrderDetailsRequest{
							{
								ProductID: 1,
								Quantity:  2,
							},
						},
					},
				},
				wantErr: true,
			},
			{
				name: "update product quantity error",
				args: args{
					resp: resp{
						productRepo: productRepo{
							getByID: getByID{
								domain: &domainProduct.Product{
									ID:       1,
									Quantity: 10,
								},
							},
							updateByQuantity: updateByQuantity{
								err: errors.New(constant.ErrorDatabaseProblem),
							},
						},
						orderRepo: orderRepo{
							save: save{
								id: 12,
							},
						},
					},
					req: &dto.CreateOrderRequest{
						OrderDetails: []dto.OrderDetailsRequest{
							{
								ProductID: 1,
								Quantity:  2,
							},
						},
					},
				},
				wantErr: true,
			},
			{
				name: "create order success",
				args: args{
					resp: resp{
						productRepo: productRepo{
							getByID: getByID{
								domain: &domainProduct.Product{
									ID:       1,
									Quantity: 10,
								},
							},
						},
						orderRepo: orderRepo{
							save: save{
								id: 12,
							},
						},
					},
					req: &dto.CreateOrderRequest{
						OrderDetails: []dto.OrderDetailsRequest{
							{
								ProductID: 1,
								Quantity:  2,
							},
						},
					},
				},
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := new(mocksPsql.ProductRepositoryIFaceMock)
			productRepo.On("GetByID", mock.Anything).Return(tt.args.resp.productRepo.getByID.domain, tt.args.resp.productRepo.getByID.err)
			productRepo.On("UpdateByQuantity", mock.Anything, mock.Anything).Return(tt.args.resp.productRepo.updateByQuantity.err)

			orderRepo := new(mocksPsql.OrderRepositoryIFaceMock)
			orderRepo.On("Save", mock.Anything).Return(tt.args.resp.orderRepo.save.id, tt.args.resp.orderRepo.save.err)

			orderDetailsRepo := new(mocksPsql.OrderDetailsRepositoryIFaceMock)
			orderDetailsRepo.On("SaveBulk", mock.Anything, mock.Anything).Return(tt.args.resp.orderDetailsRepo.saveBulk.err)

			s := usecaseOrder.NewOrderService(productRepo, orderRepo, orderDetailsRepo)

			if err := s.Save(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
