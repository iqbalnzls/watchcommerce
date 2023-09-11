package graph

import (
	"github.com/iqbalnzls/watchcommerce/src/pkg/validator"
	"github.com/iqbalnzls/watchcommerce/src/usecase/product"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	productService product.ProductServiceIFace
	v              *validator.DataValidator
}
