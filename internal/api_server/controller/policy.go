package controller

import (
	"l-iam/internal/api_server/dao"
	"l-iam/internal/api_server/service"
	v1 "l-iam/internal/pkg/model/api_server/v1"
	"l-iam/pkg/code"
	"l-iam/pkg/core"
	"l-iam/pkg/errors"

	"github.com/gin-gonic/gin"
)

type PolicyCtrl struct {
	srv service.Service
}

func NewPolicyCtrl(dao dao.Factory) *PolicyCtrl {
	return &PolicyCtrl{
		srv: service.NewService(dao),
	}
}

func (p PolicyCtrl) Create(ctx *gin.Context) {
	var policy v1.Policy
	if err := core.Parse(ctx, nil, nil, &policy); err != nil {
		core.ResponseErr(ctx, errors.WrapError(code.ErrBind, err))
		return
	}
	if err := policy.Validate(); err != nil {
		core.ResponseErr(ctx, errors.WrapError(code.ErrValidation, err))
		return
	}
	if err := p.srv.Policies().Create(ctx.Request.Context(), &policy); err != nil {
		core.ResponseErr(ctx, err)
		return
	}

	core.ResponseJson(ctx, policy)
}

func (p PolicyCtrl) List(ctx *gin.Context) {
	var opts v1.PolicyQueryOptions
	if err := core.Parse(ctx, nil, &opts, nil); err != nil {
		core.ResponseErr(ctx, errors.WrapError(code.ErrBind, err))
		return
	}
	result, err := p.srv.Policies().List(ctx, &opts)
	if err != nil {
		core.ResponseErr(ctx, err)
		return
	}

	core.ResponseJson(ctx, result)
}

func (p PolicyCtrl) Get(ctx *gin.Context) {
	var uri v1.PolicyQueryUri
	if err := core.Parse(ctx, &uri, nil, nil); err != nil {
		core.ResponseErr(ctx, errors.WrapError(code.ErrBind, err))
		return
	}

	policy, err := p.srv.Policies().Get(ctx, uri.PolicyId)
	if err != nil {
		core.ResponseErr(ctx, err)
		return
	}

	core.ResponseJson(ctx, policy)
}

func (p PolicyCtrl) Update(ctx *gin.Context) {
	var (
		uri    v1.PolicyQueryUri
		policy v1.Policy
	)

	if err := core.Parse(ctx, &uri, nil, &policy); err != nil {
		core.ResponseErr(ctx, errors.WrapError(code.ErrBind, err))
		return
	}

	if err := p.srv.Policies().Update(ctx, uri.PolicyId, &policy); err != nil {
		core.ResponseErr(ctx, err)
		return
	}

	core.ResponseJson(ctx, policy)
}

func (p PolicyCtrl) Delete(ctx *gin.Context) {
	var uri v1.PolicyQueryUri
	if err := core.Parse(ctx, &uri, nil, nil); err != nil {
		core.ResponseErr(ctx, errors.WrapError(code.ErrBind, err))
		return
	}

	if err := p.srv.Policies().Delete(ctx, uri.PolicyId); err != nil {
		core.ResponseErr(ctx, err)
		return
	}

	core.ResponseJson(ctx, nil)
}
