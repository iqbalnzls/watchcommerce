package graph

import (
	"github.com/iqbalnzls/watchcommerce/src/delivery/graph/model"
	"github.com/iqbalnzls/watchcommerce/src/dto"
)

func toProduct(d dto.GetProductResponse) *model.Product {
	return &model.Product{
		ID:        int(d.ID),
		Name:      d.Name,
		BrandName: d.BrandName,
		Price:     int(d.Price),
		Quantity:  int(d.Quantity),
	}
}

func toGetProductByBrandIDRequest(brandId int) *dto.GetProductByBrandIDRequest {
	return &dto.GetProductByBrandIDRequest{
		BrandID: int64(brandId),
	}
}

func toCreateBrandRequest(name string) *dto.CreateBrandRequest {
	return &dto.CreateBrandRequest{
		Name: name,
	}
}

func toGetProductRequest(productId int64) *dto.GetProductRequest {
	return &dto.GetProductRequest{
		ProductID: productId,
	}
}

func toProductsBrand(slice []dto.GetProductResponse) []*model.Product {
	result := make([]*model.Product, 0)

	for _, v := range slice {
		result = append(result, toProduct(v))
	}

	return result
}

func toBrand(name string) *model.Brand {
	return &model.Brand{
		Name: name,
	}
}

func toCreateProductRequest(in model.ProductInput) *dto.CreateProductRequest {
	return &dto.CreateProductRequest{
		BrandID:  int64(in.BrandID),
		Name:     in.Name,
		Price:    int64(in.Price),
		Quantity: int64(in.Quantity),
	}
}
