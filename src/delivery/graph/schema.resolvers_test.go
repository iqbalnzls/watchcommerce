package graph_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/iqbalnzls/watchcommerce/src/delivery/graph"
	"github.com/iqbalnzls/watchcommerce/src/delivery/graph/model"
	"github.com/iqbalnzls/watchcommerce/src/dto"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
	"github.com/iqbalnzls/watchcommerce/src/shared/logger"
	brandUcMock "github.com/iqbalnzls/watchcommerce/src/shared/mock/usecase/brand"
	productUcMock "github.com/iqbalnzls/watchcommerce/src/shared/mock/usecase/product"
	"github.com/iqbalnzls/watchcommerce/src/shared/validator"
)

var ctx = context.WithValue(context.Background(), constant.AppContext, appContext.NewAppContext(&logger.Log{
	XID: "3bec11b2-94a4-4f82-a1b7-229716decfd9",
}))

func Test_mutationResolver_CreateBrand(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name     string
		args     args
		wantResp *model.Brand
		wantErr  bool
		mockSave error
	}{
		{
			name: "Successful creation",
			args: args{
				name: "TestBrand",
			},
			wantResp: &model.Brand{Name: "TestBrand"},
			wantErr:  false,
		},
		{
			name: "Validation error",
			args: args{
				name: "",
			},
			wantResp: nil,
			wantErr:  true,
		},
		{
			name: "Save error",
			args: args{
				name: "TestBrand",
			},
			wantResp: nil,
			wantErr:  true,
			mockSave: errors.New("error saving brand"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			brandSvc := &brandUcMock.BrandServiceIFaceMock{}
			brandSvc.On("Save", mock.Anything, mock.Anything).Return(tt.mockSave)

			r := graph.SetupResolver(&productUcMock.ProductServiceIFaceMock{}, brandSvc, validator.NewValidator())

			gotResp, err := r.Mutation().CreateBrand(ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBrand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CreateBrand() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

//func Test_mutationResolver_CreateProduct(t *testing.T) {
//	type args struct {
//		input model.ProductInput
//	}
//	tests := []struct {
//		name     string
//		args     args
//		wantResp *model.Product
//		wantErr  bool
//		mockSave error
//	}{
//		{
//			name: "Successful creation",
//			args: args{
//				input: model.ProductInput{
//					BrandID:  1,
//					Name:     "TestProduct",
//					Price:    100,
//					Quantity: 10,
//				},
//			},
//			wantResp: &model.Product{
//				Name:     "TestProduct",
//				Price:    100,
//				Quantity: 10,
//			},
//			wantErr: false,
//		},
//		{
//			name: "Validation error",
//			args: args{
//				input: model.ProductInput{
//					Name:     "",
//					Price:    0,
//					Quantity: 0,
//				},
//			},
//			wantResp: nil,
//			wantErr:  true,
//		},
//		{
//			name: "Save error",
//			args: args{
//				input: model.ProductInput{
//					Name:     "TestProduct",
//					Price:    100,
//					Quantity: 10,
//				},
//			},
//			wantResp: nil,
//			wantErr:  true,
//			mockSave: errors.New("error saving product"),
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			productSvc := &productUcMock.ProductServiceIFaceMock{}
//			productSvc.On("Save", mock.Anything, mock.Anything).Return(tt.mockSave)
//
//			r := graph.SetupResolver(productSvc, nil, validator.NewValidator())
//
//			gotResp, err := r.Mutation().CreateProduct(ctx, tt.args.input)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotResp, tt.wantResp) {
//				t.Errorf("CreateProduct() gotResp = %v, want %v", gotResp, tt.wantResp)
//			}
//		})
//	}
//}

func Test_mutationResolver_CreateProduct(t *testing.T) {
	type args struct {
		input model.ProductInput
	}
	tests := []struct {
		name     string
		args     args
		wantResp *model.Product
		wantErr  bool
		mockSave error
	}{
		{
			name: "Success - Create product with valid input",
			args: args{
				input: model.ProductInput{
					BrandID:  1,
					Name:     "Rolex Submariner",
					Price:    10000,
					Quantity: 5,
				},
			},
			wantResp: &model.Product{
				Name:     "Rolex Submariner",
				Price:    10000,
				Quantity: 5,
			},
			wantErr:  false,
			mockSave: nil,
		},
		{
			name: "Error - Empty product name",
			args: args{
				input: model.ProductInput{
					BrandID:  1,
					Name:     "",
					Price:    10000,
					Quantity: 5,
				},
			},
			wantResp: nil,
			wantErr:  true,
			mockSave: nil,
		},
		{
			name: "Error - Zero price",
			args: args{
				input: model.ProductInput{
					BrandID:  1,
					Name:     "Rolex Submariner",
					Price:    0,
					Quantity: 5,
				},
			},
			wantResp: nil,
			wantErr:  true,
			mockSave: nil,
		},
		{
			name: "Error - Database save error",
			args: args{
				input: model.ProductInput{
					BrandID:  1,
					Name:     "Rolex Submariner",
					Price:    10000,
					Quantity: 5,
				},
			},
			wantResp: nil,
			wantErr:  true,
			mockSave: errors.New("database connection error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productSvc := &productUcMock.ProductServiceIFaceMock{}
			productSvc.On("Save", mock.Anything, mock.Anything).Return(tt.mockSave)

			r := graph.SetupResolver(productSvc, nil, validator.NewValidator())

			gotResp, err := r.Mutation().CreateProduct(ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CreateProduct() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_queryResolver_GetProduct(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name     string
		args     args
		wantResp *model.Product
		wantErr  bool
		mockGet  error
		mockData dto.GetProductResponse
	}{
		{
			name: "Successful get product",
			args: args{
				id: 1,
			},
			mockData: dto.GetProductResponse{
				ID:       1,
				Name:     "TestProduct",
				Price:    100,
				Quantity: 10,
			},
			wantResp: &model.Product{
				ID:       1,
				Name:     "TestProduct",
				Price:    100,
				Quantity: 10,
			},
			wantErr: false,
		},
		{
			name:     "Invalid ID",
			wantResp: nil,
			wantErr:  true,
		},
		{
			name: "Product not found",
			args: args{
				id: 999,
			},
			mockGet:  errors.New("product not found"),
			wantResp: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productSvc := &productUcMock.ProductServiceIFaceMock{}
			productSvc.On("Get", mock.Anything, mock.Anything).Return(tt.mockData, tt.mockGet)

			r := graph.SetupResolver(productSvc, nil, validator.NewValidator())

			gotResp, err := r.Query().GetProduct(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetProduct() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_queryResolver_GetProductsByBrandID(t *testing.T) {
	type args struct {
		brandID int
	}
	tests := []struct {
		name     string
		args     args
		wantResp []*model.Product
		wantErr  bool
		mockGet  error
		mockData []dto.GetProductResponse
	}{
		{
			name: "Successful get products by brand ID",
			args: args{
				brandID: 1,
			},
			mockData: []dto.GetProductResponse{
				{
					ID:       1,
					Name:     "Product1",
					Price:    100,
					Quantity: 10,
				},
				{
					ID:       2,
					Name:     "Product2",
					Price:    200,
					Quantity: 20,
				},
			},
			wantResp: []*model.Product{
				{
					ID:       1,
					Name:     "Product1",
					Price:    100,
					Quantity: 10,
				},
				{
					ID:       2,
					Name:     "Product2",
					Price:    200,
					Quantity: 20,
				},
			},
			wantErr: false,
		},
		{
			name:     "Invalid brand ID",
			wantResp: nil,
			wantErr:  true,
		},
		{
			name: "No products found",
			args: args{
				brandID: 999,
			},
			mockGet:  errors.New("no products found"),
			mockData: nil,
			wantResp: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productSvc := &productUcMock.ProductServiceIFaceMock{}
			productSvc.On("GetByBrandID", mock.Anything, mock.Anything).Return(tt.mockData, tt.mockGet)

			r := graph.SetupResolver(productSvc, nil, validator.NewValidator())

			gotResp, err := r.Query().GetProductsByBrandID(ctx, tt.args.brandID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProductsByBrandID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetProductsByBrandID() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
