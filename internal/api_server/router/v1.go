package router

import (
	"l-iam/internal/api_server/controller"
	"l-iam/internal/api_server/dao"
	"l-iam/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitV1Router(g *gin.Engine) {

	v1 := g.Group("v1")

	policyCtrl := controller.NewPolicyController(dao.Client())
	policy := v1.Group("policies", middleware.Publish())

	policy.GET("", nil)
	policy.POST("", policyCtrl.Create)
	policy.GET(":name", nil)
	policy.PUT(":name", nil)
	policy.DELETE(":name", nil)
}
