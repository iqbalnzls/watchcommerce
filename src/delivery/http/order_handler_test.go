package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	inHttp "github.com/iqbalnzls/watchcommerce/src/delivery/http"
	"github.com/iqbalnzls/watchcommerce/src/dto"
	"github.com/iqbalnzls/watchcommerce/src/pkg/constant"
	mocksUsecaseOrder "github.com/iqbalnzls/watchcommerce/src/pkg/mock/usecase/order"
	"github.com/iqbalnzls/watchcommerce/src/pkg/validator"
	"github.com/iqbalnzls/watchcommerce/src/usecase/order"
)

func TestNewOrderHandler(t *testing.T) {
	type args struct {
		orderService order.OrderServiceIFace
		v            *validator.DataValidator
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "order service is nil",
			args: args{
				v: validator.NewValidator(),
			},
			wantPanic: true,
		},
		{
			name: "validator is nil",
			args: args{
				orderService: new(mocksUsecaseOrder.OrderServiceIFaceMock),
			},
			wantPanic: true,
		},
		{
			name: "init order handler success",
			args: args{
				orderService: new(mocksUsecaseOrder.OrderServiceIFaceMock),
				v:            validator.NewValidator(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = inHttp.NewOrderHandler(tt.args.orderService, tt.args.v)
				})
			} else {
				assert.NotPanics(t, func() {
					_ = inHttp.NewOrderHandler(tt.args.orderService, tt.args.v)
				})
			}
		})
	}
}

func Test_orderHandler_Get(t *testing.T) {
	type orderService struct {
		data dto.GetOrderResponse
		err  error
	}
	type resp struct {
		orderService orderService
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
							"id": "test",
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
						orderService: orderService{
							err: errors.New(constant.ErrorDataNotFound),
						},
					},
				},
			},
			{
				name: "get order success",
				args: args{
					req: req{
						method: http.MethodGet,
						queryParam: map[string]string{
							"id": "1",
						},
					},
					resp: resp{
						orderService: orderService{
							data: dto.GetOrderResponse{
								ID:    1,
								Total: 10,
								Details: []dto.OrderDetailsResponse{
									{
										ID:        1,
										OrderID:   1,
										ProductID: 12,
										Quantity:  1,
										Price:     1000,
									},
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
			orderService := new(mocksUsecaseOrder.OrderServiceIFaceMock)
			orderService.On("Get", mock.Anything).Return(tt.args.resp.orderService.data, tt.args.resp.orderService.err)

			req := httptest.NewRequest(tt.args.req.method, "/api/v1/order/get", nil)
			q := req.URL.Query()
			for k, v := range tt.args.req.queryParam {
				q.Add(k, v)
			}

			req.URL.RawQuery = q.Encode()

			rec := httptest.NewRecorder()

			h := inHttp.NewOrderHandler(orderService, validator.NewValidator())
			h.Get(rec, req)
		})
	}
}

func Test_orderHandler_Save(t *testing.T) {
	type save struct {
		err error
	}
	type orderService struct {
		save save
	}
	type resp struct {
		orderService orderService
	}
	type req struct {
		method  string
		payload interface{}
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
				name: "invalid http method",
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
				name: "validate request failed",
				args: args{
					req: req{
						method: http.MethodPost,
						payload: dto.CreateOrderRequest{
							OrderDetails: []dto.OrderDetailsRequest{
								{
									ProductID: 1,
								},
								{
									Quantity: 3,
								},
							},
						},
					},
				},
			},
			{
				name: "call service error",
				args: args{
					req: req{
						method: http.MethodPost,
						payload: dto.CreateOrderRequest{
							OrderDetails: []dto.OrderDetailsRequest{
								{
									ProductID: 1,
									Quantity:  2,
								},
							},
						},
					},
					resp: resp{
						orderService: orderService{
							save: save{
								err: errors.New(constant.ErrorDatabaseProblem),
							},
						},
					},
				},
			},
			{
				name: "create order success",
				args: args{
					req: req{
						method: http.MethodPost,
						payload: dto.CreateOrderRequest{
							OrderDetails: []dto.OrderDetailsRequest{
								{
									ProductID: 1,
									Quantity:  2,
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
			orderService := new(mocksUsecaseOrder.OrderServiceIFaceMock)
			orderService.On("Save", mock.Anything).Return(tt.args.resp.orderService.save.err)

			b, _ := json.Marshal(tt.args.req.payload)

			req := httptest.NewRequest(tt.args.req.method, "/api/v1/order/get", bytes.NewReader(b))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			h := inHttp.NewOrderHandler(orderService, validator.NewValidator())
			h.Save(rec, req)
		})
	}
}
