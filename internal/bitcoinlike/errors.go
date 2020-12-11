package bitcoinlike

import "errors"

// ErrUnsupportedTypes - found unsupported types
var ErrUnsupportedTypes = errors.New("found unsupported types")

// ErrEmptyEnabledTypes - no enalbed address types
var ErrEmptyEnabledTypes = errors.New("empty enabled types")
