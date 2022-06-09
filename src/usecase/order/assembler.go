package order

import (
	domainOrder "github.com/iqbalnzls/watchcommerce/src/domain/order"
	domainOrderDetails "github.com/iqbalnzls/watchcommerce/src/domain/order_details"
	"github.com/iqbalnzls/watchcommerce/src/dto"
)

func toGetOrderResponse(order *domainOrder.Order, orderDetails []*domainOrderDetails.OrderDetails) dto.GetOrderResponse {
	return dto.GetOrderResponse{
		ID:      order.ID,
		Total:   order.Total,
		Details: toOrderDetailsResponses(orderDetails),
	}
}

func toOrderDetailsResponses(d []*domainOrderDetails.OrderDetails) []dto.OrderDetailsResponse {
	var orderDetailResponses = make([]dto.OrderDetailsResponse, 0)
	for _, v := range d {
		orderDetailResponses = append(orderDetailResponses, toOrderDetailsResponse(v))
	}

	return orderDetailResponses
}

func toOrderDetailsResponse(d *domainOrderDetails.OrderDetails) dto.OrderDetailsResponse {
	return dto.OrderDetailsResponse{
		ID:        d.ID,
		OrderID:   d.OrderID,
		ProductID: d.ProductID,
		Quantity:  d.Quantity,
		Price:     d.Price,
	}
}
