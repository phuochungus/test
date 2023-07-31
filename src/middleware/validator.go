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

func ValidatePathParam[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params T
		if err := ctx.ShouldBindUri(&params); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		ctx.Set("params", params)
		ctx.Next()
	}
}
