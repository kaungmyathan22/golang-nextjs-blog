package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/logger"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		clientIP := ctx.ClientIP()
		logger.Info(fmt.Sprintf("[%s] %s %s %v", clientIP, ctx.Request.Method, ctx.Request.RequestURI, time.Since(start)))
		ctx.Next()
	}
}
