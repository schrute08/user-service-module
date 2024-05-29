package errors

import(
	"errors"
)

var (
	ErrInvalidID = errors.New("error: invalid ID(s)")
	ErrIDNotFound = errors.New("error: ID not found")
	ErrInvalidFields = errors.New("error: invalid field(s)")
)