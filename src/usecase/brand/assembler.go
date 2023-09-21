package brand

import (
	"github.com/iqbalnzls/watchcommerce/src/domain"
	"github.com/iqbalnzls/watchcommerce/src/dto"
)

func toBrandDomain(d *dto.CreateBrandRequest) *domain.Brand {
	return &domain.Brand{
		Name: d.Name,
	}
}
