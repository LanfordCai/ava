package bitcoin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mainnetP2PKHAddresses = []string{
	"1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF",
	"1NQhfGtWRwU6zg5G58TfQibHyJEuo6ZYXw",
	"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854",
}

var mainnetP2SHAddresses = []string{
	"3NJHZpnnk3bFxqVHVS2vUomUBznju6W8D9",
	"3DtsukMi6SYqWLvE1hh5rJnHePvD77Rsga",
	"3EktnHQD7RiAE6uzMj2ZifT9YgRrkSgzQX",
}

var testnetP2PKHAddresses = []string{
	"mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
	"mrDbAZMsWY4disHVThaieUBLA29ocvM19P",
	"mx27DTNdKZgJbLHwtBJt1mcRPcejRNUMkD",
}

var testnetP2SHAddresses = []string{
	"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN",
	"2MzQwSSnBHWHqSAqtTVQ6v47XtaisrJa1Vc",
	"2NDhzMt2D9ZxXapbuq567WGeWP7NuDN81cg",
}

var invalidAddresses = []string{
	"1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF1",
	"1NQhfItWRwU6zg5G58TfQibHyJEuo6ZYXw",
	"0NT5SNNaoAXzvxRUvYGxiif93q7o9u4854",
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
