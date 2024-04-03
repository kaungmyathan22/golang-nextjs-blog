package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-blog/cmd/handlers"
)

func AuthRoutes(v1Routes *gin.RouterGroup) {
	handlers := &handlers.AuthHandler{}
	router := v1Routes.Group("/authentication")
	// public routes
	router.POST("/login", handlers.LoginHandler)
	router.POST("/register", handlers.RegisterHandler)
	// private routes
}
