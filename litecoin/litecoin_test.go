package litecoin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mainnetP2PKHAddresses = []string{
	"LUd28DgwFmae8thFhaPYHGVmps4XockqfQ",
	"LfzMXaeTd1ZvThWdTHvAiucgLBy461DdCf",
	"LY1cKc26dhtGg6EJutjSH1QXbXPbRNdVZ8",
}

var mainnetP2SHAddresses = []string{
	"MLkMXXwmuFFA5L6DmTejSgcxwZfT99f7pg",
	"MAnkEjwEDDpB6CWDeRLVJ5tD3Es2cuKALT",
	"MTf4tP1TCNBn8dNkyxeBVoPrFCcVzxJvvh",
}

var testnetP2PKHAddresses = []string{
	"mhKYHWXLmQb93RuzyVmTK2KmM8KJPFrzdb",
	"muAGGFMnSCes3hyixTDw1abTSrQDqpbVsP",
	"mvvuJvego4AeD3gqk665JVF2YEfNs35Wvd",
}

var testnetP2SHAddresses = []string{
	"Qcjgixb2zRFpZYdrp3FHPpkwYgiFdo1XaU",
	"QWjY6TjBVND7k5eEvo1EQ8Y6K5TXTvWfxN",
	"QPhKqyyS1HXZmRo31568anJPMHuzDLRBLz",
}

var invalidAddresses = []string{
	"Lhqgn2XKNtyAvGauTLhawxwLQFavDz4KYK1",
	"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854",
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
