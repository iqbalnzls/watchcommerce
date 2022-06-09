package dto

type CreateBrandRequest struct {
	Name string `json:"name" validate:"required"`
}
