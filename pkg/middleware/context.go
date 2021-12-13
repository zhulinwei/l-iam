package middleware

import (
	"l-iam/pkg/log"

	"github.com/gin-gonic/gin"
)

const (
	KeyUserName = "username"
)

func Context() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(log.KeyUsername, ctx.GetHeader(KeyUserName))
		ctx.Set(log.KeyXRequestID, ctx.GetString(XRequestIDKey))
	}
}
