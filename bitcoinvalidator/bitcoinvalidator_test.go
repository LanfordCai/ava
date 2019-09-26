package bitcoinvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = map[string][]string{
	"mainnetP2PKH": []string{
		"1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF",
		"1NQhfGtWRwU6zg5G58TfQibHyJEuo6ZYXw",
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854",
	},
	"mainnetP2SH": []string{
		"3NJHZpnnk3bFxqVHVS2vUomUBznju6W8D9",
		"3DtsukMi6SYqWLvE1hh5rJnHePvD77Rsga",
		"3EktnHQD7RiAE6uzMj2ZifT9YgRrkSgzQX",
	},
	"testnetP2PKH": []string{
		"mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
		"mrDbAZMsWY4disHVThaieUBLA29ocvM19P",
		"mx27DTNdKZgJbLHwtBJt1mcRPcejRNUMkD",
	},
	"testnetP2SH": []string{
		"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN",
		"2MzQwSSnBHWHqSAqtTVQ6v47XtaisrJa1Vc",
		"2NDhzMt2D9ZxXapbuq567WGeWP7NuDN81cg",
	},
	"invalid": []string{
		"1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF1",
		"1NQhfItWRwU6zg5G58TfQibHyJEuo6ZYXw",
		"0NT5SNNaoAXzvxRUvYGxiif93q7o9u4854",
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

func validateAddresses(t *testing.T, validator *BitcoinValidator, addresses []string, isTestnet bool, expect bool) {
	for _, a := range addresses {
		isValid := validator.IsValidAddress(a, isTestnet)
		assert.Equal(t, expect, isValid)
	}
}
