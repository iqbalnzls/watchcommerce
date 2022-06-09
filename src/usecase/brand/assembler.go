package brand

import (
	"github.com/iqbalnzls/watchcommerce/src/domain/brand"
	"github.com/iqbalnzls/watchcommerce/src/dto"
)

func toBrandDomain(d *dto.CreateBrandRequest) *brand.Brand {
	return &brand.Brand{
		Name: d.Name,
	}
}
