package router

import (
	"l-iam/internal/api_server/router/admin"
	"l-iam/internal/api_server/router/api"
	"l-iam/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	g.Use(gin.Recovery()).
		Use(middleware.Cors())
	api.InitV1(g.Group("api"))
	admin.InitV1(g.Group("admin"))
}
