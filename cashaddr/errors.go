package cashaddr

import "errors"

// List of errors
var (
	ErrInvalidCashAddr   = errors.New("[cashaddr] invalid cash addr")
	ErrInvalidLegacyAddr = errors.New("[cashaddr] invalid legacy addr")
	ErrUnsupported       = errors.New("[cashaddr] unsupported")
)
