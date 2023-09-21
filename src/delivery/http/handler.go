package http

import (
	"net/http"

	"github.com/iqbalnzls/watchcommerce/src/delivery"
)

type (
	BrandHandlerIFace interface {
		Save(w http.ResponseWriter, r *http.Request)
	}

	ProductHandlerIFace interface {
		Save(w http.ResponseWriter, r *http.Request)
		Get(w http.ResponseWriter, r *http.Request)
		GetByBrandID(w http.ResponseWriter, r *http.Request)
	}

	OrderHandlerIFace interface {
		Save(w http.ResponseWriter, r *http.Request)
		Get(w http.ResponseWriter, r *http.Request)
	}
)

type Handler struct {
	brandHandler   BrandHandlerIFace
	productHandler ProductHandlerIFace
	orderHandler   OrderHandlerIFace
}

func SetupHandler(container *delivery.Container) *Handler {
	return &Handler{
		brandHandler:   NewBrandHandler(container.BrandService, container.Validator),
		productHandler: NewProductHandler(container.ProductService, container.Validator),
		orderHandler:   NewOrderHandler(container.OrderService, container.Validator),
	}
}
