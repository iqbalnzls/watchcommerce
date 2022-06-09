package dto

type CreateProductRequest struct {
	BrandID  int64  `json:"brandID" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Price    int64  `json:"price" validate:"required"`
	Quantity int64  `json:"quantity" validate:"required"`
}

type GetProductRequest struct {
	ProductID int64 `json:"productID" validate:"required"`
}

type GetProductResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	BrandID  int64  `json:"brandID"`
	Price    int64  `json:"price"`
	Quantity int64  `json:"quantity"`
}

type GetProductByBrandIDRequest struct {
	BrandID int64 `json:"brandID" validate:"required"`
}
