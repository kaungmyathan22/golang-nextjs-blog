package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-blog/cmd/database"
)

func InitRoutes(r *gin.Engine, db *database.Queries) {
	v1Routes := r.Group("/api/v1")
	// setup auth routes
	v1Routes.GET("/monitoring/ping", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, map[string]any{
			"message": "pong",
		})
		return
	})
	AuthRoutes(v1Routes, db)
}
