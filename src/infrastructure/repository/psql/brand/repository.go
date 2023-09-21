package brand

import (
	"github.com/iqbalnzls/watchcommerce/src/domain"
)

type BrandRepositoryIFace interface {
	Save(domain *domain.Brand) (err error)
}
