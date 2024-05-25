package interfaces

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context) error
	Register(ctx *gin.Context) error
	ChangePassword(ctx *gin.Context) error
	ForgotPassword(ctx *gin.Context) error
	ResetPassword(ctx *gin.Context) error
	RefreshToken(ctx *gin.Context) error
	Logout(ctx *gin.Context) error
}

type PostHandler interface {
	CreatePost(ctx *gin.Context) error
	UpdatePost(ctx *gin.Context) error
	DeletePost(ctx *gin.Context) error
	GetPosts(ctx *gin.Context) error
	GetPost(ctx *gin.Context) error
}
