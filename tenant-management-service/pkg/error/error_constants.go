package error

import "errors"

var (
	ErrNotFound       = errors.New("record not found")
	ErrInvalidInput   = errors.New("invalid input")
	ErrUnauthorized   = errors.New("unauthorized access")
	ErrInternalServer = errors.New("internal server error")
	ErrConflict       = errors.New("conflict: resource already exists")
	ErrForbidden      = errors.New("forbidden access")
)
