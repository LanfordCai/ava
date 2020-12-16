package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLitecoinValidateAddress(t *testing.T) {
	validator := &Litecoin{}

	var mainnetCases = map[string]*Result{
		"LUd28DgwFmae8thFhaPYHGVmps4XockqfQ": {Success, true, P2PKH, ""},
		"LfzMXaeTd1ZvThWdTHvAiucgLBy461DdCf": {Success, true, P2PKH, ""},
		"LY1cKc26dhtGg6EJutjSH1QXbXPbRNdVZ8": {Success, true, P2PKH, ""},
		"MLkMXXwmuFFA5L6DmTejSgcxwZfT99f7pg": {Success, true, P2SH, ""},
		"MAnkEjwEDDpB6CWDeRLVJ5tD3Es2cuKALT": {Success, true, P2SH, ""},
		"MTf4tP1TCNBn8dNkyxeBVoPrFCcVzxJvvh": {Success, true, P2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"mhKYHWXLmQb93RuzyVmTK2KmM8KJPFrzdb": {Success, true, P2PKH, ""},
		"muAGGFMnSCes3hyixTDw1abTSrQDqpbVsP": {Success, true, P2PKH, ""},
		"mvvuJvego4AeD3gqk665JVF2YEfNs35Wvd": {Success, true, P2PKH, ""},
		"Qcjgixb2zRFpZYdrp3FHPpkwYgiFdo1XaU": {Success, true, P2SH, ""},
		"QWjY6TjBVND7k5eEvo1EQ8Y6K5TXTvWfxN": {Success, true, P2SH, ""},
		"QPhKqyyS1HXZmRo31568anJPMHuzDLRBLz": {Success, true, P2SH, ""},
	}

	for addr, result := range testnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"Lhqgn2XKNtyAvGauTLhawxwLQFavDz4KYK1":                            {Success, false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854":                             {Success, false, Unknown, ""},
		"1NQhfItWRwU6zg5G58TfQibHyJEuo6ZYXw":                             {Success, false, Unknown, ""},
		"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":                            {Success, false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ":                            {Success, false, Unknown, ""},
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P":                             {Success, false, Unknown, ""},
		"bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvze":  {Success, false, Unknown, ""},
		"bc1qf2epzuxpm32t4g02m9ya2a3lcphqg8kzp8mhchgjedg2w4n4300s0057u5": {Success, false, Unknown, ""},
		"bb1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c":                     {Success, false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":                     {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
