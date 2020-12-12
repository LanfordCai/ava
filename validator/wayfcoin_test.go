package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWayfcoinValidateAddress(t *testing.T) {
	validator := &Wayfcoin{}

	var mainnetCases = map[string]*Result{
		"WaaKv3fUm2s3H4qwpDTVdHNwjhxPoLSQFH": {true, P2PKH, ""},
		"WcPLF1mvbjNvYaS12kjaK9Gz5Gg9ZHyMGu": {true, P2PKH, ""},
		"WfdWw5VoKTgxAYYZdsMHvaJWYCy2n9gido": {true, P2PKH, ""},
		"WgxcsZERSMwTT37dUMgngt7FoyUsJuh7sL": {true, P2PKH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"DFwtUPqF3ornMxgs6gt6A3Rpcuwizzsnee":         {false, Unknown, ""},
		"WbxcsZERSMwTT37dUMgngt7FoyUsJuh7sL":         {false, Unknown, ""},
		"QN3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":        {false, Unknown, ""},
		"UcehDN7bgkJ8VgCsM96r2W2UpRoiHSQizY ":        {false, Unknown, ""},
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P":         {false, Unknown, ""},
		"WgxcsZErSMwTT37dUMgngt7FoyUsJuh7sL":         {false, Unknown, ""},
		"aajt5HHvkLTT4qwhZA1z3zP2ib2tmFm56u":         {false, Unknown, ""},
		"bb1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c": {false, Unknown, ""},
		"abcde": {false, Unknown, ""},
		"":      {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
