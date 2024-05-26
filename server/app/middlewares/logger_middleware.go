package middlewares

import (
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// start := time.Now()
		// clientIP := ctx.ClientIP()
		// logger.Info(fmt.Sprintf("[%s] %s %s %v", clientIP, ctx.Request.Method, ctx.Request.RequestURI, time.Since(start)))
		ctx.Next()
	}
}
