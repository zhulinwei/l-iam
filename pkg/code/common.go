package code

import (
	"l-iam/pkg/errors"
	"net/http"
)

// Common code
const (
	Success int = iota + 100001
	ErrUnknown
	ErrBind
	ErrValidation
	ErrTokenInvalid
	ErrPageNotFound
)

func init() {
	errors.Register(ErrUnknown, http.StatusInternalServerError, "Internal server error")
	errors.Register(ErrBind, http.StatusBadRequest, "Error occurred while binding the request body to the struct")
	errors.Register(ErrValidation, http.StatusBadRequest, "Validation failed")
	errors.Register(ErrTokenInvalid, http.StatusUnauthorized, "Token invalid")
	errors.Register(ErrPageNotFound, http.StatusForbidden, "Page not found")
}
