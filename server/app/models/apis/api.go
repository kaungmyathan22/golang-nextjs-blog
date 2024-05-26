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
	Data: map[string]string{
		"message": "Something went wrong",
	},
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

func GetStatusBadRequestResponse(data any) *APIResponse {
	return &APIResponse{
		Message: "error",
		Status:  http.StatusBadRequest,
		Data: map[string]any{
			"error": data,
		},
	}
}

func GetUnauthorizedResponse(data any) *APIResponse {
	return &APIResponse{
		Message: "error",
		Status:  http.StatusUnauthorized,
		Data: map[string]any{
			"error": data,
		},
	}
}

func GetSuccessResponse(data any) *APIResponse {
	return &APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    data,
	}
}
