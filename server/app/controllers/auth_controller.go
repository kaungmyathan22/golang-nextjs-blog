package controllers

import "github.com/gin-gonic/gin"

type AuthControllerImpl struct{}

func NewAuthControllerImpl() *AuthControllerImpl {
	return &AuthControllerImpl{}
}

func (ctrl *AuthControllerImpl) Login(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "Login",
	})
}

func (ctrl *AuthControllerImpl) Register(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "Register",
	})
}

func (ctrl *AuthControllerImpl) ChangePassword(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "ChangePassword",
	})
}

func (ctrl *AuthControllerImpl) ForgotPassword(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "ForgotPassword",
	})
}

func (ctrl *AuthControllerImpl) ResetPassword(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "ResetPassword",
	})
}

func (ctrl *AuthControllerImpl) RefreshToken(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "RefreshToken",
	})
}
