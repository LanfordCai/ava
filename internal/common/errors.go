package common

import "errors"

// ErrEmptyEnabledTypes - no enalbed address types
var ErrEmptyEnabledTypes = errors.New("Empty enabled types")

// ErrUnsupportedChain - the give chain is unsupported yet
var ErrUnsupportedChain = errors.New("Given chain is unsupported yet")

// ErrUnsupportedTypes - found unsupported types
var ErrUnsupportedTypes = errors.New("Found unsupported types")
