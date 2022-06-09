package order

import "github.com/iqbalnzls/watchcommerce/src/dto"

type OrderServiceIFace interface {
	Save(req *dto.CreateOrderRequest) (err error)
	Get(req *dto.GetOrderRequest) (resp dto.GetOrderResponse, err error)
}
