package order

import (
	domainOrder "github.com/iqbalnzls/watchcommerce/src/domain"
	"github.com/iqbalnzls/watchcommerce/src/dto"
)

func toGetOrderResponse(order *domainOrder.Order, orderDetails []*domainOrder.OrderDetails) dto.GetOrderResponse {
	return dto.GetOrderResponse{
		ID:      order.ID,
		Total:   order.Total,
		Details: toOrderDetailsResponses(orderDetails),
	}
}

func toOrderDetailsResponses(d []*domainOrder.OrderDetails) []dto.OrderDetailsResponse {
	var orderDetailResponses = make([]dto.OrderDetailsResponse, 0)
	for _, v := range d {
		orderDetailResponses = append(orderDetailResponses, toOrderDetailsResponse(v))
	}

	return orderDetailResponses
}

func toOrderDetailsResponse(d *domainOrder.OrderDetails) dto.OrderDetailsResponse {
	return dto.OrderDetailsResponse{
		ID:        d.ID,
		OrderID:   d.OrderID,
		ProductID: d.ProductID,
		Quantity:  d.Quantity,
		Price:     d.Price,
	}
}
