package litecoinvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = map[string][]string{
	"mainnetP2PKH": []string{
		"LUd28DgwFmae8thFhaPYHGVmps4XockqfQ",
		"LfzMXaeTd1ZvThWdTHvAiucgLBy461DdCf",
		"LY1cKc26dhtGg6EJutjSH1QXbXPbRNdVZ8",
	},
	"mainnetP2SH": []string{
		"MLkMXXwmuFFA5L6DmTejSgcxwZfT99f7pg",
		"MAnkEjwEDDpB6CWDeRLVJ5tD3Es2cuKALT",
		"MTf4tP1TCNBn8dNkyxeBVoPrFCcVzxJvvh",
	},
	"testnetP2PKH": []string{
		"mhKYHWXLmQb93RuzyVmTK2KmM8KJPFrzdb",
		"muAGGFMnSCes3hyixTDw1abTSrQDqpbVsP",
		"mvvuJvego4AeD3gqk665JVF2YEfNs35Wvd",
	},
	"testnetP2SH": []string{
		"Qcjgixb2zRFpZYdrp3FHPpkwYgiFdo1XaU",
		"QWjY6TjBVND7k5eEvo1EQ8Y6K5TXTvWfxN",
		"QPhKqyyS1HXZmRo31568anJPMHuzDLRBLz",
	},
	"invalid": []string{
		"Lhqgn2XKNtyAvGauTLhawxwLQFavDz4KYK1",
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854",
		"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN",
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ",
		"abcde",
		"",
	},
}

func TestIsValidAddress(t *testing.T) {
	var validator = New([]string{"P2PKH", "P2SH"})
	var validatorP2PKH = New([]string{"P2PKH"})
	var validatorP2SH = New([]string{"P2SH"})
	var validatorEmpty = New([]string{})
	for k, v := range cases {
		switch k {
		case "mainnetP2PKH":
			validateAddresses(t, validator, v, false, true)
			validateAddresses(t, validatorP2PKH, v, false, true)
			validateAddresses(t, validatorP2SH, v, false, false)
			validateAddresses(t, validatorEmpty, v, false, false)
			validateAddresses(t, validator, v, true, false)
		case "testnetP2PKH":
			validateAddresses(t, validator, v, true, true)
			validateAddresses(t, validatorP2PKH, v, true, true)
			validateAddresses(t, validatorP2SH, v, true, false)
			validateAddresses(t, validatorEmpty, v, true, false)
			validateAddresses(t, validator, v, false, false)
		case "mainnetP2SH":
			validateAddresses(t, validator, v, false, true)
			validateAddresses(t, validatorP2PKH, v, false, false)
			validateAddresses(t, validatorP2SH, v, false, true)
			validateAddresses(t, validatorEmpty, v, false, false)
			validateAddresses(t, validator, v, true, false)
		case "testnetP2SH":
			validateAddresses(t, validator, v, true, true)
			validateAddresses(t, validatorP2PKH, v, true, false)
			validateAddresses(t, validatorP2SH, v, true, true)
			validateAddresses(t, validatorEmpty, v, true, false)
			validateAddresses(t, validator, v, false, false)
		case "invalid":
			validateAddresses(t, validator, v, true, false)
			validateAddresses(t, validatorP2PKH, v, true, false)
			validateAddresses(t, validatorP2SH, v, true, false)
			validateAddresses(t, validatorEmpty, v, true, false)
			validateAddresses(t, validator, v, false, false)
		}
	}
}

func validateAddresses(t *testing.T, validator *LitecoinValidator, addresses []string, isTestnet bool, expect bool) {
	for _, a := range addresses {
		isValid := validator.IsValidAddress(a, isTestnet)
		assert.Equal(t, expect, isValid)
	}
}
