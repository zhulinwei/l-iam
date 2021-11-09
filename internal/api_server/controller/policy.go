package controller

import (
	"fmt"
	"l-iam/internal/api_server/dao"
	"l-iam/internal/api_server/model"
	"l-iam/internal/api_server/service"
	"l-iam/pkg/core"

	"github.com/gin-gonic/gin"
)

type PolicyController struct {
	service service.Service
}

func NewPolicyController(dao dao.Factory) *PolicyController {
	return &PolicyController{
		service: service.NewService(dao),
	}
}

func (p *PolicyController) Create(ctx *gin.Context) {
	var policy *model.Policy
	if err := core.Parse(ctx, nil, nil, policy); err != nil {
		core.ResponseTODO(ctx)
		return
	}
	if err := policy.Validate(); err != nil {
		core.ResponseTODO(ctx)
		return
	}
	fmt.Println(policy)
	if err := p.service.Policies().Create(ctx.Request.Context(), policy); err != nil {
		core.ResponseTODO(ctx)
		return
	}
}
