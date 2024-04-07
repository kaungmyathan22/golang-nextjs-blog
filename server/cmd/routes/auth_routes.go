package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-blog/cmd/database"
	models "github.com/kaungmyathan22/golang-blog/cmd/dto"
	"github.com/kaungmyathan22/golang-blog/cmd/handlers"
	"github.com/kaungmyathan22/golang-blog/cmd/middlewares"
)

func AuthRoutes(v1Routes *gin.RouterGroup, db *database.Queries) {
	handlers := &handlers.AuthHandler{
		Repository: db,
	}
	router := v1Routes.Group("/authentication")
	// public routes
	router.POST("/login", middlewares.ValidateRequest(&models.LoginPayload{}), handlers.LoginHandler)
	router.POST("/register", middlewares.ValidateRequest(&models.RegisterPayload{}), handlers.RegisterHandler)
	// private routes
}
