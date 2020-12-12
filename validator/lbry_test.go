package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLbryValidateAddress(t *testing.T) {
	validator := &Lbry{}

	var mainnetCases = map[string]*Result{
		"baigz7693A1aKkC6YCwgdFKM9rNZT9yWA4": {true, P2PKH, ""},
		"bajt5HHvkLTT4qwhZA1z3zP2ib2tmFm56u": {true, P2PKH, ""},
		"bakpp3dB7tWm48UjCo7U6azYejndijaghz": {true, P2PKH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"36T5s5B1Saja9yAJWxQN8SNXvNzxaoFdcy":         {false, Unknown, ""},
		"DFwtUPqF3ornMxgs6gt6A3Rpcuwizzsnee":         {false, Unknown, ""},
		"1NQhfItWRwU6zg5G58TfQibHyJEuo6ZYXw":         {false, Unknown, ""},
		"QN3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":        {false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ":        {false, Unknown, ""},
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P":         {false, Unknown, ""},
		"bakpp3d37tWm48UjCo7U6azYejndijaghz":         {false, Unknown, ""},
		"aajt5HHvkLTT4qwhZA1z3zP2ib2tmFm56u":         {false, Unknown, ""},
		"bb1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c": {false, Unknown, ""},
		"abcde": {false, Unknown, ""},
		"":      {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
