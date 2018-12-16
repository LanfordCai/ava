package qtum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mainnetP2PKHAddresses = []string{
	"QhzgLgC9DYxn9sTXWfcsFdKbFbbx2DKDL4",
	"QfFE2wjWdqPDX4TQESSdEyb5y9DqL5wYHA",
	"QNG4mkK4thjgMaFHEStfW7gseWaVhcy6QY",
}

var mainnetP2SHAddresses = []string{
	"MLkMXXwmuFFA5L6DmTejSgcxwZfT99f7pg",
	"MAnkEjwEDDpB6CWDeRLVJ5tD3Es2cuKALT",
	"M8Sk3adVUMyFJVm94JZhsCWvv7bvRyHXri",
}

var testnetP2PKHAddresses = []string{
	"qdvMzSaMH17gtpJLu33ug1cTegC5rshhg2",
	"qbgHcqxXYHVJZXHheGpHwLJsB5epDUtWxe",
	"qUcwykb7UstkWB2cLtaGjY19bfUjoxLaZS",
}

var testnetP2SHAddresses = []string{}

var invalidAddresses = []string{
	"qdvMzSaMH17gtpJLu33ug1cTegC5rshhg21",
	"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854",
	"QNG4mkK4thjgMaFHEStfW7gseWaVhcy6 QY",
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
