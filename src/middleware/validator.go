package middleware

import (
	"encoding/json"
	"net/http"
	"root/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrMsg struct {
	Value string `json:"value"`
	Tag   string `json:"tag"`
}

func ValidateBody[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body T
		if err := ctx.ShouldBind(&body); err == nil {
			ctx.Set("body", body)
			ctx.Next()
			return
		} else {
			if validationErrs, ok := err.(validator.ValidationErrors); ok {
				msges := utils.ParseErrors(validationErrs)
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"messages": msges,
				})
				return
			} else if marshallingErr, ok := err.(*json.UnmarshalTypeError); ok {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, marshallingErr.Type.String())
			}
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
