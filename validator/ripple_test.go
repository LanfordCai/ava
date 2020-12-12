package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRippleValidateAddress(t *testing.T) {
	validator := &Ripple{}

	var validCases = map[string]*Result{
		"r3kmLJN5D28dHuH8vZNUZpMC43pEHpaocV": {true, Normal, ""},
		"rKLpjpCoXgLQQYQyj13zgay73rsgmzNH13": {true, Normal, ""},
		"rPFLkxQk6xUGdGYEykqe7PR25Gr7mLHDc8": {true, Normal, ""},
		"rLwh9iqbPZ52Lg7wBNrUqoAJdddr3hs1MA": {true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"rLwh9iqbPZ52Lg7wBNrUqoAJdddr3hs1M1 ": {false, Unknown, ""},
		"1PFLkxQk6xUGdGYEykqe7PR25Gr7mLHDc8":  {false, Unknown, ""},
		"rKLpjpCoXgLQQYQ1j13zgay73rsgmzNH13":  {false, Unknown, ""},
		"abcde":                               {false, Unknown, ""},
		"":                                    {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}
}
