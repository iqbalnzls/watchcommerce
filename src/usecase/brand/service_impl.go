package brand

import (
	"github.com/iqbalnzls/watchcommerce/src/dto"
	"github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/brand"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
)

type service struct {
	brandRepo brand.RepositoryIFace
}

func NewBrandService(brandRepo brand.RepositoryIFace) ServiceIFace {
	if brandRepo == nil {
		panic("brand repository is nil")
	}

	return &service{
		brandRepo: brandRepo,
	}
}

func (s *service) Save(appCtx *appContext.AppContext, req *dto.CreateBrandRequest) (err error) {
	return s.brandRepo.Save(appCtx, toBrandDomain(req))
}
