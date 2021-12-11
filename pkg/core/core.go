package core

import (
	"l-iam/pkg/code"
	"l-iam/pkg/errors"

	"github.com/gin-gonic/gin"
)

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DataResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Parse(ctx *gin.Context, uri, query, body interface{}) error {
	if uri != nil {
		if err := ctx.ShouldBindUri(uri); err != nil {
			return err
		}
	}
	if query != nil {
		if err := ctx.ShouldBindQuery(query); err != nil {
			return err
		}
	}
	if body != nil {
		if err := ctx.ShouldBindJSON(body); err != nil {
			return err
		}
	}
	return nil
}

func ResponseErr(ctx *gin.Context, err error) {
	coder := errors.ParseErrCode(err)
	ctx.JSON(coder.HttpStatus(), ErrResponse{
		Code:    coder.Code(),
		Message: coder.Message(),
	})
}

func ResponseJson(ctx *gin.Context, data interface{}) {
	coder := code.SuccessCoder
	ctx.JSON(coder.HttpStatus(), DataResponse{
		Data:    data,
		Code:    coder.Code(),
		Message: coder.Message(),
	})
}
