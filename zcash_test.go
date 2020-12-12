package ava

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZcashValidateAddress(t *testing.T) {
	validator := &Zcash{}

	var mainnetCases = map[string]*Result{
		"t1fuvbxcLNhiPGnR2fcy4iMrntLn2y6zbiG": {true, P2PKH, ""},
		"t1J1YojHoLb5L9pvRi1sCSNgyReHrK4EbDw": {true, P2PKH, ""},
		"t1dYGZ3aCtFtXTYh42ZcyZGenUCabmysCQX": {true, P2PKH, ""},
		"t3Rqonuzz7afkF7156ZA4vi4iimRSEn41hj": {true, P2SH, ""},
		"t3fJZ5jYsyxDtvNrWBeoMbvJaQCj4JJgbgX": {true, P2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"tmVvvKBBYRnbK9JMf7AzWatJrkGh2LQwRJ7": {true, P2PKH, ""},
		"tmMjR9pDM2HLkhKzDvPd4wRoG5rAdJadrxB": {true, P2PKH, ""},
		"tmU2PyLY1hwSyZh4wi1iXBjT5pZsms68oxz": {true, P2PKH, ""},
		"t2BS7Mrbaef3fA4xrmkvDisFVXVrRBnZ6Qj": {true, P2SH, ""},
	}

	for addr, result := range testnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"t1fuvbxcLNhiPGnR2fcy4iMrntLn2y6z3iG":                            {false, Unknown, ""},
		"tmMjR9pDM3HLkhKzDvPd4wRoG5rAdJadrxB":                            {false, Unknown, ""},
		"1NQhfItWRwU6zg5G58TfQibHyJEuo6ZYXw":                             {false, Unknown, ""},
		"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":                            {false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ":                            {false, Unknown, ""},
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P":                             {false, Unknown, ""},
		"bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvze":  {false, Unknown, ""},
		"bc1qf2epzuxpm32t4g02m9ya2a3lcphqg8kzp8mhchgjedg2w4n4300s0057u5": {false, Unknown, ""},
		"bb1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c":                     {false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":                     {false, Unknown, ""},
		"abcde": {false, Unknown, ""},
		"":      {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
