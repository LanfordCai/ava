package bytom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = map[string][]string{
	"mainnetP2WPKH": []string{
		"BM1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7K23GYYF",
		"bm1qw508d6qejxtdg4y5r3zarvary0c5xw7k23gyyf",
	},
	"mainnetP2WSH": []string{
		"bm1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qk5egtg",
	},
	"testnetP2WPKH": []string{
		"tm1qw508d6qejxtdg4y5r3zarvary0c5xw7kw8fqyc",
	},
	"testnetP2WSH": []string{
		"tm1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qqq379v",
	},
	"invalid": []string{
		"bm1pw508d6qejxtdg4y5r3zarvary0c5xw7kw508d6qejxtdg4y5r3zarvary0c5xw7k7grplx",
		"BM1SW50QA3JX3S",
		"bm1zw508d6qejxtdg4y5r3zarvaryvg6kdaj",
		"tc1qw508d6qejxtdg4y5r3zarvary0c5xw7kg3g4ty ",
		"bm1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t5",
		"BM13W508D6QEJXTDG4Y5R3ZARVARY0C5XW7KN40WF2",
		"bm10w508d6qejxtdg4y5r3zarvary0c5xw7kw508d6qejxtdg4y5r3zarvary0c5xw7kw5rljs90",
		"BM1QR508D6QEJXTDG4Y5R3ZARVARYV98GJ9P",
		"tm1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3q0sL5k7",
		"tm1pw508d6qejxtdg4y5r3zarqfsj6c3",
		"tm1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3pjxtptv",
		"",
	},
}

func TestIsValidAddress(t *testing.T) {
	validator, _ := New([]string{"P2WPKH", "P2WSH"})
	validatorP2WPKH, _ := New([]string{"P2WPKH"})
	validatorP2WSH, _ := New([]string{"P2WSH"})
	for k, v := range cases {
		switch k {
		case "mainnetP2WPKH":
			validateAddresses(t, validator, v, false, true)
			validateAddresses(t, validatorP2WPKH, v, false, true)
			validateAddresses(t, validatorP2WSH, v, false, false)
			validateAddresses(t, validator, v, true, false)
		case "testnetP2WPKH":
			validateAddresses(t, validator, v, true, true)
			validateAddresses(t, validatorP2WPKH, v, true, true)
			validateAddresses(t, validatorP2WSH, v, true, false)
			validateAddresses(t, validator, v, false, false)
		case "mainnetP2WSH":
			validateAddresses(t, validator, v, false, true)
			validateAddresses(t, validatorP2WPKH, v, false, false)
			validateAddresses(t, validatorP2WSH, v, false, true)
			validateAddresses(t, validator, v, true, false)
		case "testnetP2WSH":
			validateAddresses(t, validator, v, true, true)
			validateAddresses(t, validatorP2WPKH, v, true, false)
			validateAddresses(t, validatorP2WSH, v, true, true)
			validateAddresses(t, validator, v, false, false)
		case "invalid":
			validateAddresses(t, validator, v, true, false)
			validateAddresses(t, validatorP2WPKH, v, true, false)
			validateAddresses(t, validatorP2WSH, v, true, false)
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
