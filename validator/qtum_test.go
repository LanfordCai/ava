package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQtumValidateAddress(t *testing.T) {
	validator := &Qtum{}

	var mainnetCases = map[string]*Result{
		"QhzgLgC9DYxn9sTXWfcsFdKbFbbx2DKDL4": {Success, true, P2PKH, ""},
		"QfFE2wjWdqPDX4TQESSdEyb5y9DqL5wYHA": {Success, true, P2PKH, ""},
		"QNG4mkK4thjgMaFHEStfW7gseWaVhcy6QY": {Success, true, P2PKH, ""},
		"MLkMXXwmuFFA5L6DmTejSgcxwZfT99f7pg": {Success, true, P2SH, ""},
		"MAnkEjwEDDpB6CWDeRLVJ5tD3Es2cuKALT": {Success, true, P2SH, ""},
		"M8Sk3adVUMyFJVm94JZhsCWvv7bvRyHXri": {Success, true, P2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"qdvMzSaMH17gtpJLu33ug1cTegC5rshhg2": {Success, true, P2PKH, ""},
		"qbgHcqxXYHVJZXHheGpHwLJsB5epDUtWxe": {Success, true, P2PKH, ""},
		"qUcwykb7UstkWB2cLtaGjY19bfUjoxLaZS": {Success, true, P2PKH, ""},
	}

	for addr, result := range testnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"qdvMzSaMH17gtpJLu33ug1cTegC5rshhg21":                            {Success, false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854":                             {Success, false, Unknown, ""},
		"QNG4mkK4thjgMaFHEStfW7gseWaVhcy6":                               {Success, false, Unknown, ""},
		"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":                            {Success, false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ":                            {Success, false, Unknown, ""},
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P":                             {Success, false, Unknown, ""},
		"bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvze":  {Success, false, Unknown, ""},
		"bc1qf2epzuxpm32t4g02m9ya2a3lcphqg8kzp8mhchgjedg2w4n4300s0057u5": {Success, false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":                     {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
