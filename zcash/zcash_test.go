package zcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = map[string][]string{
	"mainnetP2PKH": []string{
		"t1fuvbxcLNhiPGnR2fcy4iMrntLn2y6zbiG",
		"t1J1YojHoLb5L9pvRi1sCSNgyReHrK4EbDw",
		"t1dYGZ3aCtFtXTYh42ZcyZGenUCabmysCQX",
	},
	"mainnetP2SH": []string{
		"t3Rqonuzz7afkF7156ZA4vi4iimRSEn41hj",
		"t3fJZ5jYsyxDtvNrWBeoMbvJaQCj4JJgbgX",
	},
	"testnetP2PKH": []string{
		"tmVvvKBBYRnbK9JMf7AzWatJrkGh2LQwRJ7",
		"tmMjR9pDM2HLkhKzDvPd4wRoG5rAdJadrxB",
		"tmU2PyLY1hwSyZh4wi1iXBjT5pZsms68oxz",
	},
	"testnetP2SH": []string{
		"t2BS7Mrbaef3fA4xrmkvDisFVXVrRBnZ6Qj",
	},
	"invalid": []string{
		"t1fuvbxcLNhiPGnR2fcy4iMrntLn2y6z3iG",
		"tmMjR9pDM3HLkhKzDvPd4wRoG5rAdJadrxB",
		"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN",
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ",
		"abcde",
		"",
	},
}

func TestIsValidAddress(t *testing.T) {
	validator, _ := New([]string{"P2PKH", "P2SH"})
	validatorP2PKH, _ := New([]string{"P2PKH"})
	validatorP2SH, _ := New([]string{"P2SH"})
	for k, v := range cases {
		switch k {
		case "mainnetP2PKH":
			validateAddresses(t, validator, v, false, true)
			validateAddresses(t, validatorP2PKH, v, false, true)
			validateAddresses(t, validatorP2SH, v, false, false)
			validateAddresses(t, validator, v, true, false)
		case "testnetP2PKH":
			validateAddresses(t, validator, v, true, true)
			validateAddresses(t, validatorP2PKH, v, true, true)
			validateAddresses(t, validatorP2SH, v, true, false)
			validateAddresses(t, validator, v, false, false)
		case "mainnetP2SH":
			validateAddresses(t, validator, v, false, true)
			validateAddresses(t, validatorP2PKH, v, false, false)
			validateAddresses(t, validatorP2SH, v, false, true)
			validateAddresses(t, validator, v, true, false)
		case "testnetP2SH":
			validateAddresses(t, validator, v, true, true)
			validateAddresses(t, validatorP2PKH, v, true, false)
			validateAddresses(t, validatorP2SH, v, true, true)
			validateAddresses(t, validator, v, false, false)
		case "invalid":
			validateAddresses(t, validator, v, true, false)
			validateAddresses(t, validatorP2PKH, v, true, false)
			validateAddresses(t, validatorP2SH, v, true, false)
			validateAddresses(t, validator, v, false, false)
		}
	}
}

func validateAddresses(t *testing.T, validator *Validator, addresses []string, isTestnet bool, expect bool) {
	for _, a := range addresses {
		result := validator.IsValidAddress(a, isTestnet)
		assert.Equal(t, expect, result.IsValid)
	}
}
