package dto

type CreateProductRequest struct {
	BrandID  int64  `json:"brandID" validate:"required" example:"1"`
	Name     string `json:"name" validate:"required" example:"g-shock"`
	Price    int64  `json:"price" validate:"required" example:"12000000"`
	Quantity int64  `json:"quantity" validate:"required" example:"10"`
}

type GetProductRequest struct {
	ProductID int64 `json:"productID" validate:"required" example:"2"`
}

type GetProductResponse struct {
	ID       int64  `json:"id" example:"1"`
	Name     string `json:"name" example:"daytona"`
	BrandID  int64  `json:"brandID" example:"1"`
	Price    int64  `json:"price" example:"1000"`
	Quantity int64  `json:"quantity" example:"3"`
}

type GetProductByBrandIDRequest struct {
	BrandID int64 `json:"brandID" validate:"required" example:"12"`
}
