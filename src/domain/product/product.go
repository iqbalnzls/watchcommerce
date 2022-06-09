package product

import "time"

type Product struct {
	ID        int64
	BrandID   int64
	Name      string
	Price     int64
	Quantity  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
