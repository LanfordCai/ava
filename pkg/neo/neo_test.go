package neo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validAddresses = []string{
	"AVKeLwu1uRtM7Lf7ckkqrbtkvss7jkN38N",
	"ATzoCmmsjqPHCDPmY7mxEhSTGzpSKiTwUo",
	"AetXb6U1FMcA3zNak8FuPmY33ovi4xj4wg",
	"ASax7usrW1qwVWpW3eG14mxvP7uGQUL1eM",
}

var invalidAddresses = []string{
	"ASax7usrW1qwVWpW3eG15mxvP7uGQUL1eM",
	"ATzoCmmsjqPHCDPmY7mxEhSTGzpSKiTwUo ",
	"AEzoCmmsjqPHCDPmY7mxEhSTGzpSKiTwUo",
	"abcde",
	"",
}

func TestValidateAddress(t *testing.T) {
	validator := New()
	for _, a := range validAddresses {
		isValid, _ := validator.ValidateAddress(a, false)
		assert.Equal(t, true, isValid)
	}

	for _, a := range invalidAddresses {
		isValid, _ := validator.ValidateAddress(a, false)
		assert.Equal(t, false, isValid)
	}
}
