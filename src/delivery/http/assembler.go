package http

import "github.com/iqbalnzls/watchcommerce/src/dto"

func toBaseResponse(data interface{}) dto.BaseResponse {
	return dto.BaseResponse{
		Success: true,
		Message: "success",
		Data:    data,
	}
}
