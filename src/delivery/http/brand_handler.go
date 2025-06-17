package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/iqbalnzls/watchcommerce/src/dto"
	"github.com/iqbalnzls/watchcommerce/src/shared/app_context"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
	"github.com/iqbalnzls/watchcommerce/src/shared/validator"
	serviceBrand "github.com/iqbalnzls/watchcommerce/src/usecase/brand"
)

type brandHandler struct {
	brandService serviceBrand.ServiceIFace
	*validator.DataValidator
}

func NewBrandHandler(brandService serviceBrand.ServiceIFace, v *validator.DataValidator) BrandHandlerIFace {
	if brandService == nil {
		panic("brand service is nil")
	}
	if v == nil {
		panic("validator is nil")
	}
	return &brandHandler{
		brandService:  brandService,
		DataValidator: v,
	}
}

// Save godoc
// @Summary Save brand
// @Description API for save new brand
// @Tags Brand
// @Accept json
// @Produce  json
// @Param request body dto.CreateBrandRequest true "payload for save new brand"
// @Success 200 {object} dto.BaseResponse
// @param Authorization-Swagger header string true "Authorization for swagger purpose"
// @Router /api/v1/brand/save [post]
func (h *brandHandler) Save(w http.ResponseWriter, r *http.Request) {
	var (
		req      *dto.CreateBrandRequest
		baseResp dto.BaseResponse
		err      error
		appCtx   = app_context.ParsingAppContext(r.Context())
	)

	defer func() {
		if err != nil {
			baseResp.Message = err.Error()
		}
		b, _ := json.Marshal(baseResp)
		_, _ = w.Write(b)
		appCtx.Logger.FinishedRequest(baseResp)
	}()

	if r.Method != http.MethodPost {
		err = errors.New(constant.ErrorInvalidHttpMethod)
		appCtx.Logger.Error(err.Error())
		return
	}

	startTime := appCtx.Logger.SubProcessStart("decode request start")
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		appCtx.Logger.Error(err.Error())
		err = errors.New(constant.ErrorBadRequest)
		return
	}

	appCtx.Logger.SetRequest(req)

	if err = h.Validate(req); err != nil {
		appCtx.Logger.Error(err.Error())
		err = errors.New(constant.ErrorBadRequest)
		return
	}

	appCtx.Logger.SubProcessEnd(startTime, "decode request finish")

	if err = h.brandService.Save(appCtx, req); err != nil {
		return
	}

	baseResp = toBaseResponse(nil)

	return
}
