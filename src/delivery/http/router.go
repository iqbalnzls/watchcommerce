package http

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter(mux *http.ServeMux, middleware Middleware, handler *Handler) {
	//health check
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("service is up and running..."))
	})

	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	//brand
	mux.HandleFunc("/api/v1/brand/save", middleware(handler.brandHandler.Save))

	//product
	mux.HandleFunc("/api/v1/product/save", middleware(handler.productHandler.Save))
	mux.HandleFunc("/api/v1/product/get", middleware(handler.productHandler.Get))
	mux.HandleFunc("/api/v1/product/brand/get", middleware(handler.productHandler.GetByBrandID))

	//order
	mux.HandleFunc("/api/v1/order/save", middleware(handler.orderHandler.Save))
	mux.HandleFunc("/api/v1/order/get", middleware(handler.orderHandler.Get))
}
