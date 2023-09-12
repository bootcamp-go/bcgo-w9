package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// before handler
		now := time.Now()

		ctx.Next()

		// after handler
		duration := time.Since(now)
		method := ctx.Request.Method
		url := ctx.Request.URL.String()
		code := ctx.Writer.Status()
		log.Printf("[%s] %s %d ns - Code %d", method, url, duration.Nanoseconds(), code)
	}
}