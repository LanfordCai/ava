package bitcoinlike

import "errors"

// ErrUnsupportedTypes - found unsupported types
var ErrUnsupportedTypes = errors.New("Found unsupported types")

// ErrEmptyEnabledTypes - no enalbed address types
var ErrEmptyEnabledTypes = errors.New("Empty enabled types")
