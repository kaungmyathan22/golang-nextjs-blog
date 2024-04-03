package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-blog/cmd/handlers"
)

func AuthRoutes(router *gin.RouterGroup, handler *handlers.AuthHandler) {
	router.POST("/login", handler.LoginHandler)
}
