package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/iqbalnzls/watchcommerce/src/delivery"
	inHttp "github.com/iqbalnzls/watchcommerce/src/delivery/http"
	"github.com/iqbalnzls/watchcommerce/src/dto"
	mocksUsecaseBrand "github.com/iqbalnzls/watchcommerce/src/shared/mock/usecase/brand"
	mocks "github.com/iqbalnzls/watchcommerce/src/shared/mock/usecase/order"
	mocksUsecaseProduct "github.com/iqbalnzls/watchcommerce/src/shared/mock/usecase/product"
	"github.com/iqbalnzls/watchcommerce/src/shared/validator"
)

func TestSetupRouter(t *testing.T) {
	var (
		v              = validator.NewValidator()
		brandService   = &mocksUsecaseBrand.BrandServiceIFaceMock{}
		productService = &mocksUsecaseProduct.ProductServiceIFaceMock{}
		orderService   = &mocks.OrderServiceIFaceMock{}
		container      = &delivery.Container{
			OrderService:   orderService,
			BrandService:   brandService,
			ProductService: productService,
			Validator:      v,
		}
		mux        = http.NewServeMux()
		middleware = inHttp.SetupMiddleware()
	)

	inHttp.SetupRouter(mux, middleware, inHttp.SetupHandler(container))

	t.Run("route without middleware", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		mux.ServeHTTP(rec, req)
	})

	t.Run("route with middleware", func(t *testing.T) {
		request := &dto.CreateBrandRequest{
			Name: "90ilSx2",
		}
		b, _ := json.Marshal(request)

		req, err := http.NewRequest(http.MethodPost, "/ping", bytes.NewReader(b))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		brandService.On("Save", mock.Anything).Return(nil)

		mux.ServeHTTP(rec, req)
	})
}
