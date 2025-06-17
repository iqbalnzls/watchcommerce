package product

import (
	"github.com/iqbalnzls/watchcommerce/src/dto"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
)

type ServiceIFace interface {
	Save(appCtx *appContext.AppContext, req *dto.CreateProductRequest) (err error)
	Get(appCtx *appContext.AppContext, req *dto.GetProductRequest) (resp dto.GetProductResponse, err error)
	GetByBrandID(appCtx *appContext.AppContext, req *dto.GetProductByBrandIDRequest) (resp []dto.GetProductResponse, err error)
}
