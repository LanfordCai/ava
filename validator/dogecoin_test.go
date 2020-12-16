package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDogecoinValidateAddress(t *testing.T) {
	validator := &Dogecoin{}

	var mainnetCases = map[string]*Result{
		"D8osUgxgU4J2yEhCzN5xQQ42tuuG7qBdri": {Success, true, P2PKH, ""},
		"DFwtUPqF3ornMxgs6gt6A3Rpcuwizzsnev": {Success, true, P2PKH, ""},
		"DCZHAVah2dCyD1vmjnUm6t7Aao4YwtYWkM": {Success, true, P2PKH, ""},
		"A6T5s5B1Saja9yAJWxQN8SNXvNzxaoFdcy": {Success, true, P2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"nYF8DkhE3orcq7CypP1oEQTWWcG1yYkqec":  {Success, true, P2PKH, ""},
		"nXiLZ1g1xrcLwYq1s9wVCdiewLeY7Bw1X3":  {Success, true, P2PKH, ""},
		"nj4xWxERcYpXQLZf61R5xzMDfTjv72ajJh":  {Success, true, P2PKH, ""},
		"2NDutUPDhfMYGUFPhHRy4ZRraHVrEK7odr6": {Success, true, P2SH, ""},
		"2MwU3UtzeHG59xV7J7VSoM42U7qwCJ1GSJz": {Success, true, P2SH, ""},
	}

	for addr, result := range testnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"36T5s5B1Saja9yAJWxQN8SNXvNzxaoFdcy":                             {Success, false, Unknown, ""},
		"DFwtUPqF3ornMxgs6gt6A3Rpcuwizzsnee":                             {Success, false, Unknown, ""},
		"1NQhfItWRwU6zg5G58TfQibHyJEuo6ZYXw":                             {Success, false, Unknown, ""},
		"QN3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":                            {Success, false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ":                            {Success, false, Unknown, ""},
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P":                             {Success, false, Unknown, ""},
		"bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvze":  {Success, false, Unknown, ""},
		"bc1qf2epzuxpm32t4g02m9ya2a3lcphqg8kzp8mhchgjedg2w4n4300s0057u5": {Success, false, Unknown, ""},
		"bb1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c":                     {Success, false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":                     {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
