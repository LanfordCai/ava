package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDashValidateAddress(t *testing.T) {
	validator := &Dash{}

	var mainnetCases = map[string]*Result{
		"XpMfHazu1y8bgnP8csW73LWULGKgs4XH6U": {Success, true, P2PKH, ""},
		"Xdk9TLNrELrM4jpu6m15zZvurG66sZKi6C": {Success, true, P2PKH, ""},
		"XripctRDfASrofg4BKco41Zqvm1KaCdQpk": {Success, true, P2PKH, ""},
		"7Z5BvydGVzgbX9xqWEb1JtF9TkToG8htcV": {Success, true, P2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"yYmrsYP3XYMZr1cGtha3QzmuNB1C7CfyhV": {Success, true, P2PKH, ""},
		"yQgGqVdasi5jGfweJ84HJz4qp4ac5G2gxG": {Success, true, P2PKH, ""},
		"yUr7iXCNWDpF96W2kurtgszsrjf1XjDaUh": {Success, true, P2PKH, ""},
		"8oxdi3uiMFJ8SEob6qW76YBgg5en4352Cc": {Success, true, P2SH, ""},
		"92WthWvUyvzcJokeo3tyNgNZqzYmhbHM8N": {Success, true, P2SH, ""},
	}

	for addr, result := range testnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"yYrsYP3XYMZr1cGtha3QzmuNB1C7CfyhV":   {Success, false, Unknown, ""},
		"7Z5BvydGVzgbX9xqWEb1JtF9TkToG8htcv":  {Success, false, Unknown, ""},
		"Xdk9TLNreLrM4jpu6m15zZvurG66sZKi6C":  {Success, false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ": {Success, false, Unknown, ""},
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P":  {Success, false, Unknown, ""},
		"abcde":                               {Success, false, Unknown, ""},
		"":                                    {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
