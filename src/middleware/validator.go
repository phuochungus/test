package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateBody[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body T
		if err := ctx.BindJSON(&body); err == nil {
			ctx.Set("body", body)
			ctx.Next()
			return
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}
