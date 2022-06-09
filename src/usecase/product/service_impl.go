package product

import (
	domainProduct "github.com/iqbalnzls/watchcommerce/src/domain/product"
	"github.com/iqbalnzls/watchcommerce/src/dto"
)

type productService struct {
	productRepo domainProduct.ProductRepositoryIFace
}

func NewProductService(productRepo domainProduct.ProductRepositoryIFace) ProductServiceIFace {
	if productRepo == nil {
		panic("product repository is nil")
	}

	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) Save(req *dto.CreateProductRequest) (err error) {
	return s.productRepo.Save(toProductDomain(req))
}

func (s *productService) Get(req *dto.GetProductRequest) (resp dto.GetProductResponse, err error) {
	domain, err := s.productRepo.GetByID(req.ProductID)
	if err != nil {
		return
	}

	resp = toGetProductResponse(domain)
	return
}

func (s *productService) GetByBrandID(req *dto.GetProductByBrandIDRequest) (resp []dto.GetProductResponse, err error) {
	domains, err := s.productRepo.GetByBrandID(req.BrandID)
	if err != nil {
		return
	}

	resp = toGetProductResponses(domains)
	return
}
