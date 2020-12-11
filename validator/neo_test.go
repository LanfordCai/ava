package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNeoValidateAddress(t *testing.T) {
	validator := &Neo{}

	var mainnetCases = map[string]*Result{
		"AVKeLwu1uRtM7Lf7ckkqrbtkvss7jkN38N": {true, P2PKH, ""},
		"ATzoCmmsjqPHCDPmY7mxEhSTGzpSKiTwUo": {true, P2PKH, ""},
		"AetXb6U1FMcA3zNak8FuPmY33ovi4xj4wg": {true, P2PKH, ""},
		"ASax7usrW1qwVWpW3eG14mxvP7uGQUL1eM": {true, P2PKH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"ASax7usrW1qwVWpW3eG15mxvP7uGQUL1eM":         {false, Unknown, ""},
		"ATzoCmmsjqPHCDPmY7mxEhSTGzpSKiTwUo ":        {false, Unknown, ""},
		"AEzoCmmsjqPHCDPmY7mxEhSTGzpSKiTwUo":         {false, Unknown, ""},
		"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":        {false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ":        {false, Unknown, ""},
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P":         {false, Unknown, ""},
		"bb1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c": {false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv": {false, Unknown, ""},
		"abcde": {false, Unknown, ""},
		"":      {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
