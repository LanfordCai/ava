package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitcoinCashValidateAddress(t *testing.T) {
	validator := &BitcoinCash{}

	var mainnetCases = map[string]*Result{
		"1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF":                     {Success, true, P2PKH, ""},
		"1NQhfGtWRwU6zg5G58TfQibHyJEuo6ZYXw":                     {Success, true, P2PKH, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854":                     {Success, true, P2PKH, ""},
		"3NJHZpnnk3bFxqVHVS2vUomUBznju6W8D9":                     {Success, true, P2SH, ""},
		"3DtsukMi6SYqWLvE1hh5rJnHePvD77Rsga":                     {Success, true, P2SH, ""},
		"3EktnHQD7RiAE6uzMj2ZifT9YgRrkSgzQX":                     {Success, true, P2SH, ""},
		"bitcoincash:qzy9dedpwm53cqgcr2e3z9qy788cca6y0ce8mfkezh": {Success, true, CashAddrP2PKH, ""},
		"bitcoincash:pp297ralj957hvzyyvgmfe5vtz5c2szxqcehz8lc7f": {Success, true, CashAddrP2SH, ""},
		"bitcoincash:ppm2qsznhks23z7629mms6s4cwef74vcwvn0h829pq": {Success, true, CashAddrP2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn":                 {Success, true, P2PKH, ""},
		"mrDbAZMsWY4disHVThaieUBLA29ocvM19P":                 {Success, true, P2PKH, ""},
		"mx27DTNdKZgJbLHwtBJt1mcRPcejRNUMkD":                 {Success, true, P2PKH, ""},
		"2N3WBNpL3ZVj5PwQhSTPYZdrR7QXiKttChN":                {Success, true, P2SH, ""},
		"2MzQwSSnBHWHqSAqtTVQ6v47XtaisrJa1Vc":                {Success, true, P2SH, ""},
		"2NDhzMt2D9ZxXapbuq567WGeWP7NuDN81cg":                {Success, true, P2SH, ""},
		"bchtest:qzjncxjjzga5mr6qyy7al7g46zavnar7ggmc50uj8l": {Success, true, CashAddrP2PKH, ""},
		"bchtest:qqe0az5j8gwk2uts9v2ywdv6sd32uut9hq9hc92pee": {Success, true, CashAddrP2PKH, ""},
	}

	for addr, result := range testnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF1":                            {Success, false, Unknown, ""},
		"1NQhfItWRwU6zg5G58TfQibHyJEuo6ZYXw":                             {Success, false, Unknown, ""},
		"0NT5SNNaoAXzvxRUvYGxiif93q7o9u4854":                             {Success, false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ":                            {Success, false, Unknown, ""},
		"2b1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3q0sl5k7": {Success, false, Unknown, ""},
		"tb1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3q0sl5k7": {Success, false, Unknown, ""},
		"bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvze":  {Success, false, Unknown, ""},
		"bc1qf2epzuxpm32t4g02m9ya2a3lcphqg8kzp8mhchgjedg2w4n4300s0057u5": {Success, false, Unknown, ""},
		"bb1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c":                     {Success, false, Unknown, ""},
		"bitcoincash:qZy9dedpwm53cqgcr2e3z9qy788cca6y0ce8mfkezh":         {Success, false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":                     {Success, false, Unknown, ""},
		"bc1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c":                     {Success, false, Unknown, ""},
		"bc1qxcjkl0gyffz2tz935cepgetruee7n3kcva80a0xd9wgcyz93r2pqkgkjwv": {Success, false, Unknown, ""},
		"bc1qf2epzuxpm32t4g02m9ya2a3lcphqg8kzp8khchgjedg2w4n4300s0057u5": {Success, false, Unknown, ""},
		"bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej": {Success, false, Unknown, ""},
		"tb1qp0we5epypgj4acd2c4au58045ruud2pd6heuee":                     {Success, false, Unknown, ""},
		"tb1q63svxth22j5r73rc8xth74n5uu69n7vc0um98f":                     {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
