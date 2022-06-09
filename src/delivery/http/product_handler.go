package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/iqbalnzls/watchcommerce/src/dto"
	"github.com/iqbalnzls/watchcommerce/src/pkg/constant"
	"github.com/iqbalnzls/watchcommerce/src/pkg/utils"
	"github.com/iqbalnzls/watchcommerce/src/pkg/validator"
	usecaseProduct "github.com/iqbalnzls/watchcommerce/src/usecase/product"
)

type productHandler struct {
	productService usecaseProduct.ProductServiceIFace
	*validator.DataValidator
}

func NewProductHandler(productService usecaseProduct.ProductServiceIFace, v *validator.DataValidator) *productHandler {
	if productService == nil {
		panic("product service is nil")
	}
	if v == nil {
		panic("validator is nil")
	}

	return &productHandler{
		productService: productService,
		DataValidator:  v,
	}
}

func (h *productHandler) Save(w http.ResponseWriter, r *http.Request) {
	var (
		req      *dto.CreateProductRequest
		baseResp dto.BaseResponse
		err      error
	)

	defer func() {
		if err != nil {
			baseResp.Message = err.Error()
		}
		b, _ := json.Marshal(baseResp)
		_, _ = w.Write(b)
	}()

	if r.Method != http.MethodPost {
		err = errors.New(constant.ErrorInvalidHttpMethod)
		utils.Error(err)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(err)
		err = errors.New(constant.ErrorBadRequest)
		return
	}

	if err = h.Validate(req); err != nil {
		utils.Error(err)
		err = errors.New(constant.ErrorBadRequest)
		return
	}

	if err = h.productService.Save(req); err != nil {
		return
	}

	baseResp = toBaseResponse(nil)

	return
}

func (h *productHandler) Get(w http.ResponseWriter, r *http.Request) {
	var (
		baseResp dto.BaseResponse
		err      error
	)

	defer func() {
		if err != nil {
			baseResp.Message = err.Error()
		}
		b, _ := json.Marshal(baseResp)
		_, _ = w.Write(b)
	}()

	if r.Method != http.MethodGet {
		err = errors.New(constant.ErrorInvalidHttpMethod)
		utils.Error(err)
		return
	}

	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

	req := &dto.GetProductRequest{
		ProductID: id,
	}

	if err = h.Validate(req); err != nil {
		utils.Error(err)
		err = errors.New(constant.ErrorBadRequest)
		return
	}

	data, err := h.productService.Get(req)
	if err != nil {
		return
	}

	baseResp = toBaseResponse(data)

	return
}

func (h *productHandler) GetByBrandID(w http.ResponseWriter, r *http.Request) {
	var (
		baseResp dto.BaseResponse
		err      error
	)

	defer func() {
		if err != nil {
			baseResp.Message = err.Error()
		}
		b, _ := json.Marshal(baseResp)
		_, _ = w.Write(b)
	}()

	if r.Method != http.MethodGet {
		err = errors.New(constant.ErrorInvalidHttpMethod)
		utils.Error(err)
		return
	}

	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

	req := &dto.GetProductByBrandIDRequest{
		BrandID: id,
	}

	if err = h.Validate(req); err != nil {
		utils.Error(err)
		err = errors.New(constant.ErrorBadRequest)
		return
	}

	data, err := h.productService.GetByBrandID(req)
	if err != nil {
		return
	}

	baseResp = toBaseResponse(data)

	return
}
