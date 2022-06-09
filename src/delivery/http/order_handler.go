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
	usecaseOrder "github.com/iqbalnzls/watchcommerce/src/usecase/order"
)

type orderHandler struct {
	orderService usecaseOrder.OrderServiceIFace
	*validator.DataValidator
}

func NewOrderHandler(orderService usecaseOrder.OrderServiceIFace, v *validator.DataValidator) *orderHandler {
	if orderService == nil {
		panic("order service is nil")
	}
	if v == nil {
		panic("validator is nil")
	}

	return &orderHandler{
		orderService:  orderService,
		DataValidator: v,
	}
}

func (h *orderHandler) Save(w http.ResponseWriter, r *http.Request) {
	var (
		req      *dto.CreateOrderRequest
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

	if err = h.orderService.Save(req); err != nil {
		return
	}

	baseResp = toBaseResponse(nil)

	return
}

func (h *orderHandler) Get(w http.ResponseWriter, r *http.Request) {
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

	req := &dto.GetOrderRequest{
		OrderID: id,
	}

	if err = h.Validate(req); err != nil {
		utils.Error(err)
		err = errors.New(constant.ErrorBadRequest)
		return
	}

	data, err := h.orderService.Get(req)
	if err != nil {
		return
	}

	baseResp = toBaseResponse(data)

	return
}
