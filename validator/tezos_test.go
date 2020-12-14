package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTezosValidateAddress(t *testing.T) {
	validator := &Tezos{}

	var validCases = map[string]*Result{
		"tz1aQWeRw2MJpvvurJb78nwSpK8nvNg6EoCc": {true, Normal, ""},
		"tz1hkzS6pnfnHv9KzX1nbtqXVqUkzcem8FJs": {true, Normal, ""},
		"tz1RuFpeQEp8yC4B7yX4amBiLvpbcW7HHRrS": {true, Normal, ""},
		"tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r": {true, Normal, ""},
		"tz3RB4aoyjov4KEVRbuhvQ1CKJgBJMWhaeB8": {true, Normal, ""},
		"tz2DhC4ThUEnQXt1eWcGFCkxyXMo92hP6F6o": {true, Normal, ""},
		"tz2PAC9Br2D4skziFKxpzTcnDaso531HNn6i": {true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		// don't support KT1 address now
		"KT1G1H6niMfXH2iuMbwzi77mUpJJRJ3uLkFj":   {false, Unknown, ""},
		"tz4RuFpeQEp8yC4B7yX4amBiLvpbcW7HHRrS":   {false, Unknown, ""},
		"dz3RB4aoyjov4KEVRbuhvQ1CKJgBJMWhaeB8":   {false, Unknown, ""},
		"tz3RB4aoyjov4KEVRbuhvQ1CKJgBJMWhaeB9":   {false, Unknown, ""},
		"tz2PAC9Br2D4skzifKxpzTcnDaso531HNn6i":   {false, Unknown, ""},
		"tz2PAC9Br2D4skziFKxpzTcnDaso531HNn6i12": {false, Unknown, ""},
		"abcde":                                  {false, Unknown, ""},
		"":                                       {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}
}
