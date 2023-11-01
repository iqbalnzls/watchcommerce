package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	inHttp "github.com/iqbalnzls/watchcommerce/src/delivery/http"
	"github.com/iqbalnzls/watchcommerce/src/dto"
	"github.com/iqbalnzls/watchcommerce/src/pkg/app_context"
	"github.com/iqbalnzls/watchcommerce/src/pkg/constant"
	"github.com/iqbalnzls/watchcommerce/src/pkg/logger"
	mocksUsecaseBrand "github.com/iqbalnzls/watchcommerce/src/pkg/mock/usecase/brand"
	"github.com/iqbalnzls/watchcommerce/src/pkg/validator"
	usecaseBrand "github.com/iqbalnzls/watchcommerce/src/usecase/brand"
)

func TestNewBrandHandler(t *testing.T) {
	type args struct {
		brandService usecaseBrand.BrandServiceIFace
		v            *validator.DataValidator
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "brand service is nil",
			args: args{
				v: validator.NewValidator(),
			},
			wantPanic: true,
		},
		{
			name: "validator is nil",
			args: args{
				brandService: new(mocksUsecaseBrand.BrandServiceIFaceMock),
			},
			wantPanic: true,
		},
		{
			name: "init handler success",
			args: args{
				brandService: new(mocksUsecaseBrand.BrandServiceIFaceMock),
				v:            validator.NewValidator(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = inHttp.NewBrandHandler(tt.args.brandService, tt.args.v)
				})
			} else {
				assert.NotPanics(t, func() {
					_ = inHttp.NewBrandHandler(tt.args.brandService, tt.args.v)
				})
			}
		})
	}
}

func Test_brandHandler_Save(t *testing.T) {
	type fields struct {
		brandService  usecaseBrand.BrandServiceIFace
		DataValidator *validator.DataValidator
	}
	type brandService struct {
		err error
	}
	type resp struct {
		brandService brandService
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
		v     = validator.NewValidator()
		tests = []struct {
			name   string
			fields fields
			args   args
		}{
			{
				name: "http method is invalid",
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
						method:  http.MethodPost,
						payload: dto.CreateBrandRequest{},
					},
				},
			},
			{
				name: "call service error",
				args: args{
					req: req{
						method: http.MethodPost,
						payload: dto.CreateBrandRequest{
							Name: "rolex",
						},
					},
					resp: resp{
						brandService: brandService{
							err: errors.New(constant.ErrorDataNotFound),
						},
					},
				},
			},
			{
				name: "save brand success",
				args: args{
					req: req{
						method: http.MethodPost,
						payload: dto.CreateBrandRequest{
							Name: "rolex",
						},
					},
				},
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			brandService := new(mocksUsecaseBrand.BrandServiceIFaceMock)
			brandService.On("Save", mock.Anything).Return(tt.args.resp.brandService.err)

			b, _ := json.Marshal(tt.args.req.payload)

			req := httptest.NewRequest(tt.args.req.method, "/api/v1/brand/save", bytes.NewReader(b))
			req.Header.Set("Content-Type", "application/json")

			appCtx := app_context.NewAppContext(&logger.Log{
				XID:         uuid.New().String(),
				Time:        time.Now().String(),
				Path:        req.URL.Path,
				ServiceName: constant.AppName,
				Version:     constant.AppVersion,
				Header:      req.Header,
				IP:          req.RemoteAddr,
			})

			ctx := context.WithValue(req.Context(), constant.AppContext, appCtx)

			rec := httptest.NewRecorder()

			h := inHttp.NewBrandHandler(brandService, v)
			h.Save(rec, req.WithContext(ctx))
		})
	}
}
