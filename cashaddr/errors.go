package cashaddr

import "errors"

// List of errors
var (
	ErrInvalidCashAddr   = errors.New("invalid cash addr")
	ErrInvalidLegacyAddr = errors.New("invalid legacy addr")
	ErrUnsupported       = errors.New("unsupported")
)
