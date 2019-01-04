package dogecoin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mainnetP2PKHAddresses = []string{
	"D8osUgxgU4J2yEhCzN5xQQ42tuuG7qBdri",
	"DFwtUPqF3ornMxgs6gt6A3Rpcuwizzsnev",
	"DCZHAVah2dCyD1vmjnUm6t7Aao4YwtYWkM",
}

var mainnetP2SHAddresses = []string{
	"A6T5s5B1Saja9yAJWxQN8SNXvNzxaoFdcy",
}

var testnetP2PKHAddresses = []string{
	"nYF8DkhE3orcq7CypP1oEQTWWcG1yYkqec",
	"nXiLZ1g1xrcLwYq1s9wVCdiewLeY7Bw1X3",
	"nj4xWxERcYpXQLZf61R5xzMDfTjv72ajJh",
}

var testnetP2SHAddresses = []string{
	"2NDutUPDhfMYGUFPhHRy4ZRraHVrEK7odr6",
	"2My4aAqh43o6q27mEhp5TQLTcHYJN2n4wpC",
	"2MwU3UtzeHG59xV7J7VSoM42U7qwCJ1GSJz",
}

var invalidAddresses = []string{
	"36T5s5B1Saja9yAJWxQN8SNXvNzxaoFdcy",
	"DFwtUPqF3ornMxgs6gt6A3Rpcuwizzsnee",
	"QN3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN",
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
