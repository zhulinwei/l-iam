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

type UserCtrl struct {
	srv service.Service
}

func NewUserCtrl(dao dao.Factory) *UserCtrl {
	return &UserCtrl{
		srv: service.NewService(dao),
	}
}

func (u UserCtrl) List(ctx *gin.Context) {
	var opts v1.UserQueryOptions
	if err := core.Parse(ctx, nil, &opts, nil); err != nil {
		core.ResponseErr(ctx, errors.WrapError(code.ErrBind, err))
		return
	}

	result, err := u.srv.Users().List(ctx, &opts)
	if err != nil {
		core.ResponseErr(ctx, err)
		return
	}

	core.ResponseJson(ctx, result)
}

func (u UserCtrl) Get(ctx *gin.Context) {
	var uri v1.UserQueryUri
	if err := core.Parse(ctx, &uri, nil, nil); err != nil {
		core.ResponseErr(ctx, errors.WrapError(code.ErrBind, err))
		return
	}

	result, err := u.srv.Users().Get(ctx, uri.UserId)
	if err != nil {
		core.ResponseErr(ctx, err)
		return
	}

	core.ResponseJson(ctx, result)
}

func (u UserCtrl) Create(ctx *gin.Context) {
	var user v1.User
	if err := core.Parse(ctx, nil, nil, &user); err != nil {
		core.ResponseErr(ctx, errors.WrapError(code.ErrBind, err))
		return
	}
	if err := user.Validate(); err != nil {
		core.ResponseErr(ctx, errors.WrapError(code.ErrValidation, err))
		return
	}

	if err := u.srv.Users().Create(ctx, &user); err != nil {
		core.ResponseErr(ctx, err)
		return
	}

	core.ResponseJson(ctx, user)
}

func (u UserCtrl) Update(ctx *gin.Context) {
	var (
		user v1.User
		uri  v1.UserQueryUri
	)
	if err := core.Parse(ctx, &uri, nil, &user); err != nil {
		core.ResponseErr(ctx, errors.WrapError(code.ErrBind, err))
		return
	}
	if err := user.Validate(); err != nil {
		core.ResponseErr(ctx, errors.WrapError(code.ErrValidation, err))
		return
	}
	if err := u.srv.Users().Update(ctx, uri.UserId, &user); err != nil {
		core.ResponseErr(ctx, err)
		return
	}
	core.ResponseJson(ctx, user)
}

func (u UserCtrl) Delete(ctx *gin.Context) {
	var uri v1.UserQueryUri
	if err := core.Parse(ctx, &uri, nil, nil); err != nil {
		core.ResponseErr(ctx, errors.WrapError(code.ErrBind, err))
		return
	}

	if err := u.srv.Users().Delete(ctx, uri.UserId); err != nil {
		core.ResponseErr(ctx, err)
		return
	}

	core.ResponseJson(ctx, nil)
}
