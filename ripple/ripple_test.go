package ripple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validAddresses = []string{
	"r3kmLJN5D28dHuH8vZNUZpMC43pEHpaocV",
	"rKLpjpCoXgLQQYQyj13zgay73rsgmzNH13",
	"rPFLkxQk6xUGdGYEykqe7PR25Gr7mLHDc8",
	"rLwh9iqbPZ52Lg7wBNrUqoAJdddr3hs1MA",
}

var invalidAddresses = []string{
	"rLwh9iqbPZ52Lg7wBNrUqoAJdddr3hs1M1",
	"1PFLkxQk6xUGdGYEykqe7PR25Gr7mLHDc8 ",
	"rKLpjpCoXgLQQYQ1j13zgay73rsgmzNH13",
	"abcde",
	"",
}

func TestIsValidAddress(t *testing.T) {
	validator := New()
	for _, a := range validAddresses {
		isValid := validator.IsValidAddress(a, false)
		assert.Equal(t, true, isValid)
	}

	for _, a := range invalidAddresses {
		isValid := validator.IsValidAddress(a, false)
		assert.Equal(t, false, isValid)
	}
}
