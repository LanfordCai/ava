package ava

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: testnet casese
func TestYcashValidateAddress(t *testing.T) {
	validator := &Ycash{}

	var mainnetCases = map[string]*Result{
		"s1jF4pKgr8izDWNLZ9D1AtJfWFhgCY1Qexd": {true, P2PKH, ""},
		"s1MLh26NK6cMAPQqxBbuJcKVgo1C1tBkT3j": {true, P2PKH, ""},
		"s1gsQmQeieHAMh8caW9f5jDTVqZUmTsQBsX": {true, P2PKH, ""},
		"s35qM2AnDARUhfFnWYirrcXbeTd4f7SE7T5": {true, P2SH, ""},
		"s3KJ6JzL72o2rLXdwdpW9HjqW94NHCssbq8": {true, P2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"s1MBh26NK6cMAPQqxBbuJcKVgo1C1tBkT3j":                            {false, Unknown, ""},
		"s1jF4pKgr8izDWNLZ9D1AtJfWFMgCY1Qexd":                            {false, Unknown, ""},
		"t1gsQmQeieHAMh8caW9f5jDTVqZUmTsQBsX":                            {false, Unknown, ""},
		"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":                            {false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ":                            {false, Unknown, ""},
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P":                             {false, Unknown, ""},
		"bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvze":  {false, Unknown, ""},
		"bc1qf2epzuxpm32t4g02m9ya2a3lcphqg8kzp8mhchgjedg2w4n4300s0057u5": {false, Unknown, ""},
		"bb1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c":                     {false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":                     {false, Unknown, ""},
		"abcde": {false, Unknown, ""},
		"":      {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
