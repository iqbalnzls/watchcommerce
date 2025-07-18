package product

import (
	domainProduct "github.com/iqbalnzls/watchcommerce/src/domain"
	"github.com/iqbalnzls/watchcommerce/src/dto"
)

func toProductDomain(d *dto.CreateProductRequest) *domainProduct.Product {
	return &domainProduct.Product{
		BrandID:  d.BrandID,
		Name:     d.Name,
		Price:    d.Price,
		Quantity: d.Quantity,
	}
}

func toGetProductResponse(d *domainProduct.Product) dto.GetProductResponse {
	return dto.GetProductResponse{
		ID:        d.ID,
		Name:      d.Name,
		BrandName: d.Brand.Name,
		Price:     d.Price,
		Quantity:  d.Quantity,
	}
}

func toGetProductResponses(d []*domainProduct.Product) []dto.GetProductResponse {
	var domains = make([]dto.GetProductResponse, 0)
	for _, v := range d {
		domains = append(domains, toGetProductResponse(v))
	}

	return domains
}
