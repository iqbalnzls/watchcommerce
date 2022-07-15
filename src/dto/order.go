package dto

type CreateOrderRequest struct {
	OrderDetails []OrderDetailsRequest `json:"orderDetails" validate:"dive"`
}

type OrderDetailsRequest struct {
	ProductID int64 `json:"productID" validate:"required" example:"1"`
	Quantity  int64 `json:"quantity" validate:"required" example:"12"`
}

type GetOrderRequest struct {
	OrderID int64 `json:"orderID" validate:"required" example:"23"`
}

type GetOrderResponse struct {
	ID      int64                  `json:"id" example:"11"`
	Total   int64                  `json:"total" example:"10"`
	Details []OrderDetailsResponse `json:"details"`
}

type OrderDetailsResponse struct {
	ID        int64 `json:"id" example:"1"`
	OrderID   int64 `json:"orderID" example:"12"`
	ProductID int64 `json:"productID" example:"4"`
	Quantity  int64 `json:"quantity" example:"3"`
	Price     int64 `json:"price" example:"13400990"`
}
