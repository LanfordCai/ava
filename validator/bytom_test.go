package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytomValidateAddress(t *testing.T) {
	validator := &Bytom{}

	var mainnetCases = map[string]*Result{
		"BM1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7K23GYYF":                     {Success, true, P2WPKH, ""},
		"bm1qw508d6qejxtdg4y5r3zarvary0c5xw7k23gyyf":                     {Success, true, P2WPKH, ""},
		"bm1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qk5egtg": {Success, true, P2WSH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"tm1qw508d6qejxtdg4y5r3zarvary0c5xw7kw8fqyc":                     {Success, true, P2WPKH, ""},
		"tm1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qqq379v": {Success, true, P2WSH, ""},
	}

	for addr, result := range testnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"bm1pw508d6qejxtdg4y5r3zarvary0c5xw7kw508d6qejxtdg4y5r3zarvary0c5xw7k7grplx": {Success, false, Unknown, ""},
		"BM1SW50QA3JX3S":                                                               {Success, false, Unknown, ""},
		"bm1zw508d6qejxtdg4y5r3zarvaryvg6kdaj":                                         {Success, false, Unknown, ""},
		"tc1qw508d6qejxtdg4y5r3zarvary0c5xw7kg3g4ty":                                   {Success, false, Unknown, ""},
		"bm1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t5 ":                                  {Success, false, Unknown, ""},
		"BM13W508D6QEJXTDG4Y5R3ZARVARY0C5XW7KN40WF2":                                   {Success, false, Unknown, ""},
		"bm10w508d6qejxtdg4y5r3zarvary0c5xw7kw508d6qejxtdg4y5r3zarvary0c5xw7kw5rljs90": {Success, false, Unknown, ""},
		"BM1QR508D6QEJXTDG4Y5R3ZARVARYV98GJ9P":                                         {Success, false, Unknown, ""},
		"tm1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3q0sL5k7":               {Success, false, Unknown, ""},
		"tm1pw508d6qejxtdg4y5r3zarqfsj6c3":                                             {Success, false, Unknown, ""},
		"tm1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3pjxtptv":               {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
