package models

type LoginPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResult struct {
	Token string `json:"token" binding:"required"`
}

type RegisterPayload struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordPayload struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type ForgotPasswordPayload struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordPayload struct {
	Email string `json:"email" binding:"required"`
}

type MessageResponse struct {
	Message string `json:"message" binding:"required"`
}
