package base58check

import "errors"

// List of errors
var (
	ErrInvalidChar             = errors.New("[base58check] invalid char")
	ErrInvalidChecksum         = errors.New("[base58check] invalid checksum")
	ErrUnsupportedChecksumType = errors.New("[base58check] unsupported checksum type")
)
