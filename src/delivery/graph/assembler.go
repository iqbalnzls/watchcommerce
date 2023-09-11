package graph

import (
	"github.com/iqbalnzls/watchcommerce/src/delivery/graph/model"
	"github.com/iqbalnzls/watchcommerce/src/dto"
)

func toProduct(d dto.GetProductResponse) *model.Product {
	return &model.Product{
		ID:       int(d.ID),
		Name:     d.Name,
		BrandID:  int(d.BrandID),
		Price:    int(d.Price),
		Quantity: int(d.Quantity),
	}
}
