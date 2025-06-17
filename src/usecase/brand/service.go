package brand

import (
	"github.com/iqbalnzls/watchcommerce/src/dto"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
)

type ServiceIFace interface {
	Save(appCtx *appContext.AppContext, req *dto.CreateBrandRequest) (err error)
}
