package apis

type LoginPayload struct {
	Email    string `json:"email" binding:"required" valid:"email~Invalid email address"`
	Password string `json:"password" binding:"required" valid:"sixToEightDigitAlphanumericPasswordValidator~Password must be between 6 to 8 alphanumeric characters"`
}

type RegisterPayload struct {
	Email    string `json:"email" binding:"required" valid:"email~Please provide valid email address"`
	Password string `json:"password" binding:"required" valid:"sixToEightDigitAlphanumericPasswordValidator~Password must be between 6 to 8 alphanumeric characters"`
	Name     string `json:"name" binding:"required" valid:"required~Name is required"`
}

type ChangePasswordPayload struct {
	OldPassword string `json:"oldPassword" binding:"required" valid:"sixToEightDigitAlphanumericPasswordValidator~Old Password must be between 6 to 8 alphanumeric characters"`
	NewPassword string `json:"newPassword" binding:"required" valid:"sixToEightDigitAlphanumericPasswordValidator~New Password must be between 6 to 8 alphanumeric characters"`
}

type ForgotPasswordPayload struct {
	Email string `json:"email" binding:"required" valid:"email~Please provide valid email address"`
}

type ResetPasswordPayload struct {
	Password string `json:"password" binding:"required" valid:"sixToEightDigitAlphanumericPasswordValidator~Password must be between 6 to 8 alphanumeric characters"`
}
