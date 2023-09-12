package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const TokenKey = "MY_TOKEN"

func Authenticator() gin.HandlerFunc {
	token := os.Getenv(TokenKey)
	return func(ctx *gin.Context) {
		// before handler
		tokenHeader := ctx.GetHeader("Authorization")
		if tokenHeader != token {
			code := http.StatusUnauthorized
			body := gin.H{"error": "user unauthorized"}
			ctx.JSON(code, body)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
