package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func Publish() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		var res string
		paths := strings.Split(ctx.Request.URL.Path, "/")
		if len(paths) > 2 {
			res = paths[2]
		}
		method := ctx.Request.Method

		// TODO 使用redis发送变更消息
		switch res {
		case "policies":
			notify(method)
		case "secrets":
			notify(method)
		default:
		}
	}
}

func notify(method string) {
	switch method {
	case "POST", "PUT", "DELETE", "PATH":
		// TODO
	}
}
