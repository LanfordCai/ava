package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDashValidateAddress(t *testing.T) {
	validator := &Dash{}

	var mainnetCases = map[string]*Result{
		"XpMfHazu1y8bgnP8csW73LWULGKgs4XH6U": {true, P2PKH, ""},
		"Xdk9TLNrELrM4jpu6m15zZvurG66sZKi6C": {true, P2PKH, ""},
		"XripctRDfASrofg4BKco41Zqvm1KaCdQpk": {true, P2PKH, ""},
		"7Z5BvydGVzgbX9xqWEb1JtF9TkToG8htcV": {true, P2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"yYmrsYP3XYMZr1cGtha3QzmuNB1C7CfyhV": {true, P2PKH, ""},
		"yQgGqVdasi5jGfweJ84HJz4qp4ac5G2gxG": {true, P2PKH, ""},
		"yUr7iXCNWDpF96W2kurtgszsrjf1XjDaUh": {true, P2PKH, ""},
		"8oxdi3uiMFJ8SEob6qW76YBgg5en4352Cc": {true, P2SH, ""},
		"92WthWvUyvzcJokeo3tyNgNZqzYmhbHM8N": {true, P2SH, ""},
	}

	for addr, result := range testnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"yYrsYP3XYMZr1cGtha3QzmuNB1C7CfyhV":   {false, Unknown, ""},
		"7Z5BvydGVzgbX9xqWEb1JtF9TkToG8htcv":  {false, Unknown, ""},
		"Xdk9TLNreLrM4jpu6m15zZvurG66sZKi6C":  {false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ": {false, Unknown, ""},
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P":  {false, Unknown, ""},
		"abcde":                               {false, Unknown, ""},
		"":                                    {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
