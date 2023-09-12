package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticator() gin.HandlerFunc {
	token := "123456"

	return func(ctx *gin.Context) {
		// before handler
		tokenHeader := ctx.GetHeader("Authorization")
		if tokenHeader != token {
			code := http.StatusUnauthorized
			body := gin.H{"message": "Unauthorized"}

			ctx.JSON(code, body)
			ctx.Abort()
			return
		}

		ctx.Next()

		// after handler
		// ...
	}
}