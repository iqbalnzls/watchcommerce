package dto

type CreateOrderRequest struct {
	OrderDetails []OrderDetailsRequest `json:"orderDetails,dive"`
}

type OrderDetailsRequest struct {
	ProductID int64 `json:"productID" validate:"required"`
	Quantity  int64 `json:"quantity" validate:"required"`
}

type GetOrderRequest struct {
	OrderID int64 `json:"orderID" validate:"required"`
}

type GetOrderResponse struct {
	ID      int64                  `json:"id"`
	Total   int64                  `json:"total"`
	Details []OrderDetailsResponse `json:"details"`
}

type OrderDetailsResponse struct {
	ID        int64 `json:"id"`
	OrderID   int64 `json:"orderID"`
	ProductID int64 `json:"productID"`
	Quantity  int64 `json:"quantity"`
	Price     int64 `json:"price"`
}
