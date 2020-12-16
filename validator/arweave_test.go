package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArweaveValidateAddress(t *testing.T) {
	validator := &Arweave{}

	var validCases = map[string]*Result{
		"kEGgWKgTtojj-TGdbIUPPHmntIWQEs1IPegkCH3SyZ8": {Success, true, Normal, ""},
		"181f-LjI1ooRZhj3wLK2bCjOpgUmfOo_BplG4mCGuxU": {Success, true, Normal, ""},
		"Tk1NuG7Jxr9Ecgva5tWOJya2QGDOoS6hMZP0paB129c": {Success, true, Normal, ""},
		"XepqBRwnq8SNfple3WHh0VWo06P8KQ0hVlNOjOgrJ5w": {Success, true, Normal, ""},
		"tO710xPXwCGwTPGKtOkEq3PMbWqXs9jOGiL8TCpDuw0": {Success, true, Normal, ""},
		"_m4ftvKoEnbB7toHVBkuZWXYRK0j1mmgyHsug2ayffY": {Success, true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"tO710xCGwTPGKtOkEq3PMbWqXs9jOGiL8TCpDuw0":     {Success, false, Unknown, ""},
		"_m4ftvKoEnbB7toHVBkuZWXYRK0j1mmgyHsug2ayffY1": {Success, false, Unknown, ""},
		"XepqBRwnq8SNfple3W#h0VWo06P8KQ0hVlNOjOgrJ5w":  {Success, false, Unknown, ""},
		"LjI1ooRZhj3wLK20000bCjOpgUmfOo_BplG4mCGuxU":   {Success, false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":   {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}
}
