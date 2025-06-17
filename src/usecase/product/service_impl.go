package product

import (
	"github.com/iqbalnzls/watchcommerce/src/dto"
	domainProduct "github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/product"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
)

type productService struct {
	productRepo domainProduct.RepositoryIFace
}

func NewProductService(productRepo domainProduct.RepositoryIFace) ServiceIFace {
	if productRepo == nil {
		panic("product repository is nil")
	}

	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) Save(appCtx *appContext.AppContext, req *dto.CreateProductRequest) (err error) {
	return s.productRepo.Save(appCtx, toProductDomain(req))
}

func (s *productService) Get(appCtx *appContext.AppContext, req *dto.GetProductRequest) (resp dto.GetProductResponse, err error) {
	domain, err := s.productRepo.GetByID(appCtx, req.ProductID)
	if err != nil {
		return
	}

	resp = toGetProductResponse(domain)
	return
}

func (s *productService) GetByBrandID(appCtx *appContext.AppContext, req *dto.GetProductByBrandIDRequest) (resp []dto.GetProductResponse, err error) {
	domains, err := s.productRepo.GetByBrandID(appCtx, req.BrandID)
	if err != nil {
		appCtx.Logger.Error(err.Error())
		return
	}

	resp = toGetProductResponses(domains)
	return
}
