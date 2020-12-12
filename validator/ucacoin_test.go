package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUcacoinValidateAddress(t *testing.T) {
	validator := &Ucacoin{}

	var mainnetCases = map[string]*Result{
		"UaFYCsCHXtbW4EWmckYXP8BmGgezJrRdem": {true, P2PKH, ""},
		"Uak5JfpmLyjJkGJP37eBWHVPQmYo9wLX7w": {true, P2PKH, ""},
		"UbehDN7bgkJ8VgCsM96r2W2UpRoiHSQizY": {true, P2PKH, ""},
		"UcAkUrswRtP6vwZQf2F26k68dyHqWG7Sw5": {true, P2PKH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"UcDkUrswRtP6vwZQf2F26k68dyHqWG7Sw5":         {false, Unknown, ""},
		"DFwtUPqF3ornMxgs6gt6A3Rpcuwizzsnee":         {false, Unknown, ""},
		"1NQhfItWRwU6zg5G58TfQibHyJEuo6ZYXw":         {false, Unknown, ""},
		"QN3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":        {false, Unknown, ""},
		"UcehDN7bgkJ8VgCsM96r2W2UpRoiHSQizY ":        {false, Unknown, ""},
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
