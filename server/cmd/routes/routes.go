package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	v1Routes := r.Group("/api/v1")
	// setup auth routes
	AuthRoutes(v1Routes)
}
