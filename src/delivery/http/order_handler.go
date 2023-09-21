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

func NewOrderHandler(orderService usecaseOrder.OrderServiceIFace, v *validator.DataValidator) OrderHandlerIFace {
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

// Save godoc
// @Summary Create order
// @Description API for create new order
// @Tags Order
// @Accept json
// @Produce  json
// @Param request body dto.CreateOrderRequest true "payload for create new order"
// @Success 200 {object} dto.BaseResponse
// @param Authorization-Swagger header string true "Authorization for swagger purpose"
// @Router /api/v1/order/save [post]
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

// Get godoc
// @Summary Get order by id
// @Description API for get new order by id
// @Tags Order
// @Produce  json
// @Param id query string true "order id"
// @Success 200 {object} dto.BaseResponse{data=dto.GetOrderResponse,success=bool,message=string}
// @Router /api/v1/order/get [get]
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
