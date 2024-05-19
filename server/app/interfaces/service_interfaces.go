package interfaces

type LoginPayload struct {
	Email    string
	Password string
}

type LoginResult struct {
	token string
}

type RegisterPayload struct {
	Username string
	Email    string
	Name     string
	Password string
}

type ChangePasswordPayload struct {
	OldPassword string
	NewPassword string
}

type ForgotPasswordPayload struct {
	Email string
}

type ResetPasswordPayload struct {
	Email string
}

type MessageResponse struct {
	Message string
}

type AuthService interface {
	Login(payload LoginPayload) (LoginResult, error)
	Register(payload RegisterPayload) (MessageResponse, error)
	ChangePassword(payload ChangePasswordPayload) (string, error)
	ForgotPassword(payload ForgotPasswordPayload) (string, error)
	ResetPassword(payload ResetPasswordPayload) (string, error)
	RefreshToken() (string, error)
}
