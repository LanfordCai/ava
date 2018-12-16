package zcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mainnetP2PKHAddresses = []string{
	"t1fuvbxcLNhiPGnR2fcy4iMrntLn2y6zbiG",
	"t1J1YojHoLb5L9pvRi1sCSNgyReHrK4EbDw",
	"t1dYGZ3aCtFtXTYh42ZcyZGenUCabmysCQX",
}

var mainnetP2SHAddresses = []string{
	"t3Rqonuzz7afkF7156ZA4vi4iimRSEn41hj",
	"t3fJZ5jYsyxDtvNrWBeoMbvJaQCj4JJgbgX",
}

var testnetP2PKHAddresses = []string{
	"tmVvvKBBYRnbK9JMf7AzWatJrkGh2LQwRJ7",
	"tmMjR9pDM2HLkhKzDvPd4wRoG5rAdJadrxB",
	"tmU2PyLY1hwSyZh4wi1iXBjT5pZsms68oxz",
}

var testnetP2SHAddresses = []string{
	"t2BS7Mrbaef3fA4xrmkvDisFVXVrRBnZ6Qj",
}

var invalidAddresses = []string{
	"t1fuvbxcLNhiPGnR2fcy4iMrntLn2y6z3iG",
	"tmMjR9pDM3HLkhKzDvPd4wRoG5rAdJadrxB",
	"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN",
	"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ",
	"abcde",
	"",
}

func TestIsP2PKH(t *testing.T) {
	for _, a := range mainnetP2PKHAddresses {
		isValid := IsP2PKH(a, false)
		assert.Equal(t, true, isValid)
	}

	for _, a := range testnetP2PKHAddresses {
		isValid := IsP2PKH(a, true)
		assert.Equal(t, true, isValid)
	}
}

func TestIsP2SH(t *testing.T) {
	for _, a := range mainnetP2SHAddresses {
		isValid := IsP2SH(a, false)
		assert.Equal(t, true, isValid)
	}

	for _, a := range testnetP2SHAddresses {
		isValid := IsP2SH(a, true)
		assert.Equal(t, true, isValid)
	}
}

func TestIsValidAddress(t *testing.T) {
	mainnetAddresses := append([]string(nil), mainnetP2PKHAddresses...)
	mainnetAddresses = append(mainnetAddresses, mainnetP2SHAddresses...)

	testnetAddresses := append([]string(nil), testnetP2PKHAddresses...)
	testnetAddresses = append(testnetAddresses, testnetP2SHAddresses...)

	for _, a := range mainnetAddresses {
		isValid := IsValidAddress(a, false)
		assert.Equal(t, true, isValid)
	}

	for _, a := range testnetAddresses {
		isValid := IsValidAddress(a, true)
		assert.Equal(t, true, isValid)
	}

	for _, a := range invalidAddresses {
		isValidMainnetAddr := IsValidAddress(a, false)
		isValidTestnetAddr := IsValidAddress(a, true)
		assert.Equal(t, false, isValidMainnetAddr)
		assert.Equal(t, false, isValidTestnetAddr)
	}
}
