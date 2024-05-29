package errors

import(
	"errors"
)

var (
	ErrInvalidID = errors.New("error: invalid ID(s)")
	ErrUserNotFound = errors.New("error: user not found")
	ErrInvalidFields = errors.New("error: invalid field(s)")
)