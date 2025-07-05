package graph_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/iqbalnzls/watchcommerce/src/delivery/graph"
	"github.com/iqbalnzls/watchcommerce/src/delivery/graph/model"
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
