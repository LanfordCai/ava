package common

import "errors"

// ErrEmptyEnabledTypes - no enalbed address types
var ErrEmptyEnabledTypes = errors.New("Empty enabled types")

// ErrUnsupportedTypes - found unsupported types
var ErrUnsupportedTypes = errors.New("Found unsupported types")
