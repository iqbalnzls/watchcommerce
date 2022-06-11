package brand

import "github.com/iqbalnzls/watchcommerce/src/dto"

type BrandServiceIFace interface {
	Save(req *dto.CreateBrandRequest) (err error)
}
