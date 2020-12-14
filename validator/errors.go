package validator

import "errors"

// List of Errors
var (
	ErrUnsupported     = errors.New("[validator] unsupported")
	ErrInvalidAddrType = errors.New("[validator] invalid address type")
)
