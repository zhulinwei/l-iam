package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

var _ IAuth = (*AutoStrategy)(nil)

type AutoStrategy struct {
	basic *BasicStrategy
}

func NewAutoStrategy(basic *BasicStrategy) AutoStrategy {
	return AutoStrategy{
		basic: basic,
	}
}

func (a AutoStrategy) AuthFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := Auth{}
		// 最多返回2个数组
		header := strings.SplitN(ctx.GetHeader("Authorization"), " ", 2)

		switch header[0] {
		case "Basic":
			auth.SetStrategy(a.basic)
		case "Bearer":
		default:
		}
		auth.AuthFunc()(ctx)

		ctx.Next()
	}
}
