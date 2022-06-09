package product

import "github.com/iqbalnzls/watchcommerce/src/dto"

type ProductServiceIFace interface {
	Save(req *dto.CreateProductRequest) (err error)
	Get(req *dto.GetProductRequest) (resp dto.GetProductResponse, err error)
	GetByBrandID(req *dto.GetProductByBrandIDRequest) (resp []dto.GetProductResponse, err error)
}
