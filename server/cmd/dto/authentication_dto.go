package dto

type RegisterPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=5"`
	FullName string `json:"fullName" validate:"required"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=5"`
}

type ChangePasswordPayload struct {
	CurrentPassword string `json:"currentPassword" validate:"required,gte=5"`
	NewPassword     string `json:"newPassword" validate:"required,gte=5"`
}
