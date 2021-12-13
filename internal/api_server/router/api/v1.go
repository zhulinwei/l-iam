package api

import (
	"l-iam/internal/api_server/controller"
	"l-iam/internal/api_server/dao"
	"l-iam/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitV1(g *gin.RouterGroup) {
	v1 := g.Group("v1")
	v1.Use(middleware.RequestID())

	dbClient := dao.Client()

	userCtrl := controller.NewUserCtrl(dbClient)
	user := v1.Group("users")
	user.GET("", userCtrl.List)
	user.POST("", userCtrl.Create)

	user.GET(":user_id", userCtrl.Get)
	user.PUT(":user_id", userCtrl.Update)
	user.DELETE(":user_id", userCtrl.Delete)

	policyCtrl := controller.NewPolicyCtrl(dbClient)
	policy := v1.Group("policies", middleware.Publish())

	policy.GET("", policyCtrl.List)
	policy.POST("", policyCtrl.Create)
	policy.GET(":policy_id", policyCtrl.Get)
	policy.PUT(":policy_id", policyCtrl.Update)
	policy.DELETE(":policy_id", policyCtrl.Delete)
}
