package errors

import (
	"fmt"
	"net/http"
	"sync"
)

type ErrCoder struct {
	code       int // 业务状态码
	httpStatus int // http状态码
	err        error
	message    string
}

func (e ErrCoder) Code() int {
	return e.code
}

func (e ErrCoder) SetCode(code int) {
	e.code = code
}

func (e ErrCoder) Message() string {
	if e.err != nil {
		return e.err.Error()
	}
	return e.message
}

func (e ErrCoder) SetMessage(msg string) {
	e.message = msg
}

func (e ErrCoder) HttpStatus() int {
	return e.httpStatus
}

func (e ErrCoder) SetHttpStatus(httpStatus int) {
	e.httpStatus = httpStatus
}

var UnknownErrCoder = ErrCoder{
	code:       0,
	httpStatus: http.StatusInternalServerError,
	message:    "An internal server error occurred",
}

func (e ErrCoder) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return e.message
}

func WrapError(code int, err error) error {
	return &ErrCoder{err: err, code: code}
}

func WithCode(code int, message string, args ...interface{}) error {
	return &ErrCoder{err: fmt.Errorf(message, args...), code: code}
}

func ParseErrCode(err error) ErrCoder {
	errCoder, ok := err.(*ErrCoder)
	if !ok {
		return UnknownErrCoder
	}

	registeredCoder := codes[errCoder.code]
	if errCoder.Code() <= 0 {
		errCoder.SetCode(registeredCoder.Code())
	}
	if errCoder.Message() == "" {
		errCoder.SetMessage(registeredCoder.Message())
	}
	if errCoder.HttpStatus() <= 0 {
		errCoder.SetHttpStatus(registeredCoder.HttpStatus())
	}

	return *errCoder
}

var codes = map[int]ErrCoder{}
var codeMux = &sync.Mutex{}

func Register(errCode, httpStatus int, message string) {
	coder := ErrCoder{
		code:       errCode,
		message:    message,
		httpStatus: httpStatus,
	}
	codeMux.Lock()
	defer codeMux.Unlock()

	codes[coder.Code()] = coder
}
