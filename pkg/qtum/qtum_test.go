package qtum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = map[string][]string{
	"mainnetP2PKH": []string{
		"QhzgLgC9DYxn9sTXWfcsFdKbFbbx2DKDL4",
		"QfFE2wjWdqPDX4TQESSdEyb5y9DqL5wYHA",
		"QNG4mkK4thjgMaFHEStfW7gseWaVhcy6QY",
	},
	"mainnetP2SH": []string{
		"MLkMXXwmuFFA5L6DmTejSgcxwZfT99f7pg",
		"MAnkEjwEDDpB6CWDeRLVJ5tD3Es2cuKALT",
		"M8Sk3adVUMyFJVm94JZhsCWvv7bvRyHXri",
	},
	"testnetP2PKH": []string{
		"qdvMzSaMH17gtpJLu33ug1cTegC5rshhg2",
		"qbgHcqxXYHVJZXHheGpHwLJsB5epDUtWxe",
		"qUcwykb7UstkWB2cLtaGjY19bfUjoxLaZS",
	},
	"testnetP2SH": []string{},
	"invalid": []string{
		"qdvMzSaMH17gtpJLu33ug1cTegC5rshhg21",
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854",
		"QNG4mkK4thjgMaFHEStfW7gseWaVhcy6 QY",
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ",
		"abcde",
		"",
	},
}

func TestValidateAddress(t *testing.T) {
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
		isValid, _ := validator.ValidateAddress(a, isTestnet)
		assert.Equal(t, expect, isValid)
	}
}
