package order

import (
	"github.com/iqbalnzls/watchcommerce/src/dto"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
)

type ServiceIFace interface {
	Save(appCtx *appContext.AppContext, req *dto.CreateOrderRequest) (err error)
	Get(appCtx *appContext.AppContext, req *dto.GetOrderRequest) (resp dto.GetOrderResponse, err error)
}
