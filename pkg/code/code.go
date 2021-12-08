package code

import (
	"net/http"
)

type Coder struct {
	code       int // 业务状态码
	httpStatus int // http状态码
	message    string
}

func (c Coder) Code() int {
	return c.code
}

func (c Coder) Message() string {
	return c.message
}

func (c Coder) HttpStatus() int {
	return c.httpStatus
}

var SuccessCoder = Coder{code: Success, httpStatus: http.StatusOK, message: "OK"}
