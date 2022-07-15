package dto

type BaseResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data,omitempty"`
}
