package apis

import "net/http"

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var UnauthorizedResponse = APIResponse{
	Message: "error",
	Status:  http.StatusUnauthorized,
	Data:    nil,
}

var InternalServerErrorResponse = APIResponse{
	Message: "error",
	Status:  http.StatusInternalServerError,
	Data:    nil,
}
var BadRequestResponse = APIResponse{
	Message: "error",
	Status:  http.StatusBadRequest,
	Data:    nil,
}

func GetStatusAcceptedResponse(data any) *APIResponse {
	return &APIResponse{
		Message: "success",
		Status:  http.StatusOK,
		Data:    data,
	}
}
