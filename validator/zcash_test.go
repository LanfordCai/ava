package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZcashValidateAddress(t *testing.T) {
	validator := &Zcash{}

	var mainnetCases = map[string]*Result{
		"t1fuvbxcLNhiPGnR2fcy4iMrntLn2y6zbiG": {Success, true, P2PKH, ""},
		"t1J1YojHoLb5L9pvRi1sCSNgyReHrK4EbDw": {Success, true, P2PKH, ""},
		"t1dYGZ3aCtFtXTYh42ZcyZGenUCabmysCQX": {Success, true, P2PKH, ""},
		"t3Rqonuzz7afkF7156ZA4vi4iimRSEn41hj": {Success, true, P2SH, ""},
		"t3fJZ5jYsyxDtvNrWBeoMbvJaQCj4JJgbgX": {Success, true, P2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"tmVvvKBBYRnbK9JMf7AzWatJrkGh2LQwRJ7": {Success, true, P2PKH, ""},
		"tmMjR9pDM2HLkhKzDvPd4wRoG5rAdJadrxB": {Success, true, P2PKH, ""},
		"tmU2PyLY1hwSyZh4wi1iXBjT5pZsms68oxz": {Success, true, P2PKH, ""},
		"t2BS7Mrbaef3fA4xrmkvDisFVXVrRBnZ6Qj": {Success, true, P2SH, ""},
	}

	for addr, result := range testnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"t1fuvbxcLNhiPGnR2fcy4iMrntLn2y6z3iG":                            {Success, false, Unknown, ""},
		"tmMjR9pDM3HLkhKzDvPd4wRoG5rAdJadrxB":                            {Success, false, Unknown, ""},
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
