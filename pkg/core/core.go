package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// todo
func ResponseJson(ctx *gin.Context, statusCode int, body interface{}) {
	ctx.JSON(statusCode, body)
}

func ResponseTODO(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
