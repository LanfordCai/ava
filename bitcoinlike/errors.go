package bitcoinlike

import "errors"

// ErrEmptyAcceptableTypes - no acceptable address types
var ErrEmptyAcceptableTypes = errors.New("Empty acceptable types")

// ErrUnsupportedTypes - found unsupported types
var ErrUnsupportedTypes = errors.New("Found unsupported types")
