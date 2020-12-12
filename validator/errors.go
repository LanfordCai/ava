package validator

import "errors"

// List of Errors
var (
	ErrUnsupported     = errors.New("unsupported")
	ErrInvalidAddrType = errors.New("invalid address type")
)
