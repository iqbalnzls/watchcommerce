package domain

import "time"

type OrderDetails struct {
	ID        int64
	OrderID   int64
	ProductID int64
	Quantity  int64
	Price     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
