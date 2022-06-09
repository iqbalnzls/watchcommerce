package http

import "net/http"

func SetupRouter(mux *http.ServeMux, middleware Middleware, handler *handler) {
	//brand
	mux.HandleFunc("/api/v1/brand/save", middleware(handler.brand.Save))

	//product
	mux.HandleFunc("/api/v1/product/save", middleware(handler.product.Save))
	mux.HandleFunc("/api/v1/product/get", middleware(handler.product.Get))
	mux.HandleFunc("/api/v1/product/brand/get", middleware(handler.product.GetByBrandID))

	//order
	mux.HandleFunc("/api/v1/order/save", middleware(handler.order.Save))
	mux.HandleFunc("/api/v1/order/get", middleware(handler.order.Get))
}
