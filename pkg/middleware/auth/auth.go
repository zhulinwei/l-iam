package auth

import "github.com/gin-gonic/gin"

// IAuth 认证
type IAuth interface {
	AuthFunc() gin.HandlerFunc
}

type Auth struct {
	strategy IAuth
}

func (o *Auth) SetStrategy(strategy IAuth) {
	o.strategy = strategy
}

func (o *Auth) AuthFunc() gin.HandlerFunc {
	return o.strategy.AuthFunc()
}
