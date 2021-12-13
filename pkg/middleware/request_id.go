package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	XRequestIDKey = "X-Request-ID"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.GetHeader(XRequestIDKey)

		if requestID == "" {
			requestID = uuid.New().String()
			ctx.Set(XRequestIDKey, requestID)
			ctx.Request.Header.Set(XRequestIDKey, requestID)
		}

		// 把XRequestIDKey作为响应体的头部字段返回
		ctx.Writer.Header().Set(XRequestIDKey, requestID)
		ctx.Next()
	}
}
