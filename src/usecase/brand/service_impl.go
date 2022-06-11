package brand

import (
	"github.com/iqbalnzls/watchcommerce/src/domain/brand"
	"github.com/iqbalnzls/watchcommerce/src/dto"
)

type service struct {
	brandRepo brand.BrandRepositoryIFace
}

func NewBrandService(brandRepo brand.BrandRepositoryIFace) BrandServiceIFace {
	if brandRepo == nil {
		panic("brand repository is nil")
	}

	return &service{
		brandRepo: brandRepo,
	}
}

func (s *service) Save(req *dto.CreateBrandRequest) (err error) {
	return s.brandRepo.Save(toBrandDomain(req))
}
