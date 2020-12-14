package base58check

import "errors"

// List of errors
var (
	ErrInvalidInput            = errors.New("[base58check] invalid input")
	ErrInvalidChecksum         = errors.New("[base58check] invalid checksum")
	ErrUnsupportedChecksumType = errors.New("[base58check] unsupported checksum type")
)
