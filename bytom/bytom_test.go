package bytom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mainnetP2WPKHAddresses = []string{
	"BM1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7K23GYYF",
	"bm1qw508d6qejxtdg4y5r3zarvary0c5xw7k23gyyf",
}

var mainnetP2WSHAddresses = []string{
	"bm1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qk5egtg",
}

var testnetP2WPKHAddresses = []string{
	"tm1qw508d6qejxtdg4y5r3zarvary0c5xw7kw8fqyc",
}

var testnetP2WSHAddresses = []string{
	"tm1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qqq379v",
}

var invalidAddresses = []string{
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
}

func TestIsP2WPKH(t *testing.T) {
	for _, a := range mainnetP2WPKHAddresses {
		isValid := IsP2WPKH(a, false)
		assert.Equal(t, true, isValid)
	}

	for _, a := range testnetP2WPKHAddresses {
		isValid := IsP2WPKH(a, true)
		assert.Equal(t, true, isValid)
	}
}

func TestIsP2WSH(t *testing.T) {
	for _, a := range mainnetP2WSHAddresses {
		isValid := IsP2WSH(a, false)
		assert.Equal(t, true, isValid)
	}

	for _, a := range testnetP2WSHAddresses {
		isValid := IsP2WSH(a, true)
		assert.Equal(t, true, isValid)
	}
}

func TestIsValidAddress(t *testing.T) {
	mainnetAddresses := append([]string(nil), mainnetP2WPKHAddresses...)
	mainnetAddresses = append(mainnetAddresses, mainnetP2WSHAddresses...)

	testnetAddresses := append([]string(nil), testnetP2WPKHAddresses...)
	testnetAddresses = append(testnetAddresses, testnetP2WSHAddresses...)

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
