package http

import "github.com/iqbalnzls/watchcommerce/src/delivery"

type handler struct {
	brand   *brandHandler
	product *productHandler
	order   *orderHandler
}

func SetupHandler(container *delivery.Container) *handler {
	return &handler{
		brand:   NewBrandHandler(container.BrandService, container.Validator),
		product: NewProductHandler(container.ProductService, container.Validator),
		order:   NewOrderHandler(container.OrderService, container.Validator),
	}
}
