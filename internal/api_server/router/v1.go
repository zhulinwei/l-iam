package router

import (
	"l-iam/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func initV1Router(g *gin.Engine) {

	v1 := g.Group("v1")
	policy := g.Group("policies", middleware.Publish())

	policy.GET("", nil)
	policy.POST("", nil)
	policy.GET(":name", nil)
	policy.PUT(":name", nil)
	policy.DELETE(":name", nil)
}
