package bitcoinlike

import "errors"

// ErrEmptyEnabledTypes - no acceptable address types
var ErrEmptyEnabledTypes = errors.New("Empty acceptable types")

// ErrUnsupportedTypes - found unsupported types
var ErrUnsupportedTypes = errors.New("Found unsupported types")
