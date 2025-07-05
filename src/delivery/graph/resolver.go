package graph

import (
	"github.com/iqbalnzls/watchcommerce/src/shared/validator"
	"github.com/iqbalnzls/watchcommerce/src/usecase/brand"
	"github.com/iqbalnzls/watchcommerce/src/usecase/product"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func SetupResolver(productService product.ServiceIFace, brandService brand.ServiceIFace, v *validator.DataValidator) *Resolver {
	return &Resolver{
		productService: productService,
		brandService:   brandService,
		v:              v,
	}
}

type Resolver struct {
	productService product.ServiceIFace
	brandService   brand.ServiceIFace
	v              *validator.DataValidator
}
