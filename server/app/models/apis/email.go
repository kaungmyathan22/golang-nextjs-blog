package apis

type WelcomeEmail struct {
	Name string
	To   string
}

type ForgotPasswordEmail struct {
	Email string
	Code  string
}
