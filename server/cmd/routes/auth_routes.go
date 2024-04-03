package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-blog/cmd/handlers"
	"github.com/kaungmyathan22/golang-blog/cmd/middlewares"
	"github.com/kaungmyathan22/golang-blog/cmd/models"
)

func AuthRoutes(v1Routes *gin.RouterGroup) {
	handlers := &handlers.AuthHandler{}
	router := v1Routes.Group("/authentication")
	// public routes
	router.POST("/login", middlewares.ValidateRequest(&models.LoginPayload{}), handlers.LoginHandler)
	router.POST("/register", handlers.RegisterHandler)
	// private routes
}
