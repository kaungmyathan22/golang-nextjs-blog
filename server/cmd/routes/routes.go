package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-blog/cmd/handlers"
)

func InitRoutes(r *gin.Engine) {
	v1Routes := r.Group("/v1")
	/** auth region */
	authHandlers := &handlers.AuthHandler{}
	authRouter := v1Routes.Group("/authentication")
	AuthRoutes(authRouter, authHandlers)
}
