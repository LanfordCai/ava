package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitcoinValidateAddress(t *testing.T) {
	validator := &Bitcoin{}

	var mainnetCases = map[string]*Result{
		"1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF":                             {true, P2PKH, ""},
		"1NQhfGtWRwU6zg5G58TfQibHyJEuo6ZYXw":                             {true, P2PKH, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854":                             {true, P2PKH, ""},
		"3NJHZpnnk3bFxqVHVS2vUomUBznju6W8D9":                             {true, P2SH, ""},
		"3DtsukMi6SYqWLvE1hh5rJnHePvD77Rsga":                             {true, P2SH, ""},
		"3EktnHQD7RiAE6uzMj2ZifT9YgRrkSgzQX":                             {true, P2SH, ""},
		"bc1q2l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":                     {true, P2WPKH, ""},
		"bc1q86ml6tnunc2cs30centm2dnqqammrzqhkflc98":                     {true, P2WPKH, ""},
		"bc1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c":                     {true, P2WPKH, ""},
		"bc1qxcjkl0gyffz2tz935cepgetruee7n3kcva80a0xd9wgcyz93r2pqkgkjwv": {true, P2WSH, ""},
		"bc1qf2epzuxpm32t4g02m9ya2a3lcphqg8kzp8khchgjedg2w4n4300s0057u5": {true, P2WSH, ""},
		"bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej": {true, P2WSH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn":                             {true, P2PKH, ""},
		"mrDbAZMsWY4disHVThaieUBLA29ocvM19P":                             {true, P2PKH, ""},
		"mx27DTNdKZgJbLHwtBJt1mcRPcejRNUMkD":                             {true, P2PKH, ""},
		"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":                            {true, P2SH, ""},
		"2MzQwSSnBHWHqSAqtTVQ6v47XtaisrJa1Vc":                            {true, P2SH, ""},
		"2NDhzMt2D9ZxXapbuq567WGeWP7NuDN81cg":                            {true, P2SH, ""},
		"tb1q63svxth22j5r73rc8xth74n5uu69n7vc0um98f":                     {true, P2WPKH, ""},
		"tb1qp0we5epypgj4acd2c4au58045ruud2pd6heuee":                     {true, P2WPKH, ""},
		"tb1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3q0sl5k7": {true, P2WSH, ""},
	}

	for addr, result := range testnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"1YURbVuocZZZPi8LPU6GfAcKShYY7hLXbrG75v9zBXbS2zaqaHfSmGJvNEZwU3oETNZdPNxqLwR5C": {false, Unknown, ""},
		"1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF1":                                           {false, Unknown, ""},
		"1NQhfItWRwU6zg5G58TfQibHyJEuo6ZYXw":                                            {false, Unknown, ""},
		"0NT5SNNaoAXzvxRUvYGxiif93q7o9u4854":                                            {false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ":                                           {false, Unknown, ""},
		"2b1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3q0sl5k7":                {false, Unknown, ""},
		"tb1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3q0sl5k7":                {false, Unknown, ""},
		"bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvze":                 {false, Unknown, ""},
		"bc1qf2epzuxpm32t4g02m9ya2a3lcphqg8kzp8mhchgjedg2w4n4300s0057u5":                {false, Unknown, ""},
		"bb1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c":                                    {false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":                                    {false, Unknown, ""},
		"abcde": {false, Unknown, ""},
		"":      {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
