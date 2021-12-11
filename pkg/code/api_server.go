package code

import (
	"l-iam/pkg/errors"
	"net/http"
)

const (
	ErrUserNotFount int = iota + 110001
	ErrUserAlreadyExist
)

func init() {
	errors.Register(ErrUserNotFount, http.StatusBadRequest, "User not found")
	errors.Register(ErrUserAlreadyExist, http.StatusBadRequest, "User already exist")
}
