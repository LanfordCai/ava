package sia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validAddresses = []string{
	"00322de65dca5d1683989d5f1647261393d6faf98c0e61dbb098bc9a21473995fba12fcabd9f",
	"00a60fe14a896249f15b19f962186d5e576dae8369afa53146947564f2600d1ab5086aadf6ff",
	"00e6bc7d4ca22eddab4d7a2cb3144657600aaa1890318a2908d6b3ea15ce3c39f14e732f3f95",
	"00fe5af357ef5dc4ec3d0b4d74de1c021420fbebfcd764b1bfd005695c12daff46419f79d610",
	"01507cde7ac75a643128f55dc36e3e82a5b8759343c2cdb3659b3a5a25e3560adc3d333463dc",
}

var invalidAddresses = []string{
	"00322de65dca5d1683989d5f1647261393d6faf98c0e61dbb098bc9a21473995fba12fcabd9f ",
	"00a60fe14a896249f15b19f962186d5e576dae8369afa53146947564f2600d1ab5086adf6ff",
	"00e6bc7d4ca22eddab4d7a2cb3144657600aaa1890318a2908d6b3ea15ce3c39f14e73883f95",
	"10fe5af357ef5dc4ec3d0b4d74de1c021420fbebfcd764b1bfd005695c12daff46419f79d610",
	"01507cde7ac75a643128 f55dc36e3e82a5b8759343c2cdb3659b3a5a25e3560adc3d333463dc",
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
