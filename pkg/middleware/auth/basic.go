package auth

import "github.com/gin-gonic/gin"

var _ IAuth = (*BasicStrategy)(nil)

type BasicStrategy struct {
	compare func(username, password string) bool
}

func NewBasicStrategy(compare func(username, password string) bool) BasicStrategy {
	return BasicStrategy{compare: compare}
}

func (b *BasicStrategy) AuthFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
