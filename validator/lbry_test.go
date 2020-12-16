package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLbryValidateAddress(t *testing.T) {
	validator := &Lbry{}

	var mainnetCases = map[string]*Result{
		"baigz7693A1aKkC6YCwgdFKM9rNZT9yWA4": {Success, true, P2PKH, ""},
		"bajt5HHvkLTT4qwhZA1z3zP2ib2tmFm56u": {Success, true, P2PKH, ""},
		"bakpp3dB7tWm48UjCo7U6azYejndijaghz": {Success, true, P2PKH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"36T5s5B1Saja9yAJWxQN8SNXvNzxaoFdcy":         {Success, false, Unknown, ""},
		"DFwtUPqF3ornMxgs6gt6A3Rpcuwizzsnee":         {Success, false, Unknown, ""},
		"1NQhfItWRwU6zg5G58TfQibHyJEuo6ZYXw":         {Success, false, Unknown, ""},
		"QN3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":        {Success, false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ":        {Success, false, Unknown, ""},
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P":         {Success, false, Unknown, ""},
		"bakpp3d37tWm48UjCo7U6azYejndijaghz":         {Success, false, Unknown, ""},
		"aajt5HHvkLTT4qwhZA1z3zP2ib2tmFm56u":         {Success, false, Unknown, ""},
		"bb1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c": {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
