package brand

import "github.com/iqbalnzls/watchcommerce/src/dto"

type BrandServicaIFace interface {
	Save(req *dto.CreateBrandRequest) (err error)
}
