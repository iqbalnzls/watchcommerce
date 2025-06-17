package brand

import (
	"github.com/iqbalnzls/watchcommerce/src/domain"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
)

type RepositoryIFace interface {
	Save(appCtx *appContext.AppContext, domain *domain.Brand) (err error)
}
