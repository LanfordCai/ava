package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitcoinGoldValidateAddress(t *testing.T) {
	validator := &BitcoinGold{}

	var mainnetCases = map[string]*Result{
		"GezKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P": {Success, true, P2PKH, ""},
		"GP8Kd2o6xu9qqj1K1qFuM2Ls3FBZ2Sm5tV": {Success, true, P2PKH, ""},
		"GM83B4wED9tcZEqugSimmkaPyMbvK3LEeq": {Success, true, P2PKH, ""},
		"ASac6nDf6B6RWrPfWtg15q5yx2ipwGpdik": {Success, true, P2SH, ""},
		"AK6CpKN5hrs7KVcFezvnRB3B6QQP8X6Zw2": {Success, true, P2SH, ""},
		"AahbMnQVKPJGRtudeQvg5uhVWYrpnRYDkF": {Success, true, P2SH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"moQeqV3MCPRtPC9UxEaFTFop6gqwXjzsfW":  {Success, true, P2PKH, ""},
		"mraKKEsKHn7BA6qTcZA5ZRbXnfGXSRRFGW":  {Success, true, P2PKH, ""},
		"n2Haq3wDkFx1uMUb6jbQ2xZ853oP8ybzdZ":  {Success, true, P2PKH, ""},
		"2N9tLowg62V6CFY3jkDV5B4PoeMVPtCuCyq": {Success, true, P2SH, ""},
		"2N9iShvRYeLpmQkuwwz41WCmDinkfLgg3X3": {Success, true, P2SH, ""},
		"2NAZSywv2d6Vztd1FcF11ZXHt5TqEmWnLHS": {Success, true, P2SH, ""},
	}

	for addr, result := range testnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"ASac6nDf6B6RWrPfWtg15q5yx2ipwGpdiK":                             {Success, false, Unknown, ""},
		"1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF1":                            {Success, false, Unknown, ""},
		"1NQhfItWRwU6zg5G58TfQibHyJEuo6ZYXw":                             {Success, false, Unknown, ""},
		"0NT5SNNaoAXzvxRUvYGxiif93q7o9u4854":                             {Success, false, Unknown, ""},
		"1NT5SNNaoAXzvxRUvYGxiif93q7o9u4854 ":                            {Success, false, Unknown, ""},
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P":                             {Success, false, Unknown, ""},
		"bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvze":  {Success, false, Unknown, ""},
		"bc1qf2epzuxpm32t4g02m9ya2a3lcphqg8kzp8mhchgjedg2w4n4300s0057u5": {Success, false, Unknown, ""},
		"bb1qql2qamp2az7h5ejnjyuxt4294watgcmrd76n8c":                     {Success, false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":                     {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
