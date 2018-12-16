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

func TestIsValidAddress(t *testing.T) {
	for _, a := range validAddresses {
		isValid := IsValidAddress(a)
		assert.Equal(t, true, isValid)
	}

	for _, a := range invalidAddresses {
		isValid := IsValidAddress(a)
		assert.Equal(t, false, isValid)
	}
}
