package apis

type LoginResponse struct {
	Token string `json:"token" binding:"required"`
}

type MessageResponse struct {
	Message string `json:"message" binding:"required"`
}
