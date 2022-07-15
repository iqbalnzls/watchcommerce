package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/iqbalnzls/watchcommerce/src/dto"
	"github.com/iqbalnzls/watchcommerce/src/pkg/constant"
	"github.com/iqbalnzls/watchcommerce/src/pkg/utils"
	"github.com/iqbalnzls/watchcommerce/src/pkg/validator"
	serviceBrand "github.com/iqbalnzls/watchcommerce/src/usecase/brand"
)

type brandHandler struct {
	brandService serviceBrand.BrandServiceIFace
	*validator.DataValidator
}

func NewBrandHandler(brandService serviceBrand.BrandServiceIFace, v *validator.DataValidator) *brandHandler {
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

	if err = h.brandService.Save(req); err != nil {
		return
	}

	baseResp = toBaseResponse(nil)

	return
}
