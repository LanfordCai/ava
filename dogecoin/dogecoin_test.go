package dogecoin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = map[string][]string{
	"mainnetP2PKH": []string{
		"D8osUgxgU4J2yEhCzN5xQQ42tuuG7qBdri",
		"DFwtUPqF3ornMxgs6gt6A3Rpcuwizzsnev",
		"DCZHAVah2dCyD1vmjnUm6t7Aao4YwtYWkM",
	},
	"mainnetP2SH": []string{
		"A6T5s5B1Saja9yAJWxQN8SNXvNzxaoFdcy",
	},
	"testnetP2PKH": []string{
		"nYF8DkhE3orcq7CypP1oEQTWWcG1yYkqec",
		"nXiLZ1g1xrcLwYq1s9wVCdiewLeY7Bw1X3",
		"nj4xWxERcYpXQLZf61R5xzMDfTjv72ajJh",
	},
	"testnetP2SH": []string{
		"2NDutUPDhfMYGUFPhHRy4ZRraHVrEK7odr6",
		"2My4aAqh43o6q27mEhp5TQLTcHYJN2n4wpC",
		"2MwU3UtzeHG59xV7J7VSoM42U7qwCJ1GSJz",
	},
	"invalid": []string{
		"36T5s5B1Saja9yAJWxQN8SNXvNzxaoFdcy",
		"DFwtUPqF3ornMxgs6gt6A3Rpcuwizzsnee",
		"QN3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN",
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
