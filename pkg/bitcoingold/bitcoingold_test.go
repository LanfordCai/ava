package bitcoingold

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = map[string][]string{
	"mainnetP2PKH": []string{
		"GezKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P",
		"GP8Kd2o6xu9qqj1K1qFuM2Ls3FBZ2Sm5tV",
		"GM83B4wED9tcZEqugSimmkaPyMbvK3LEeq",
	},
	"mainnetP2SH": []string{
		"ASac6nDf6B6RWrPfWtg15q5yx2ipwGpdik",
		"AK6CpKN5hrs7KVcFezvnRB3B6QQP8X6Zw2",
		"AahbMnQVKPJGRtudeQvg5uhVWYrpnRYDkF",
	},
	"testnetP2PKH": []string{
		"moQeqV3MCPRtPC9UxEaFTFop6gqwXjzsfW",
		"mraKKEsKHn7BA6qTcZA5ZRbXnfGXSRRFGW",
		"n2Haq3wDkFx1uMUb6jbQ2xZ853oP8ybzdZ",
	},
	"testnetP2SH": []string{
		"2N9tLowg62V6CFY3jkDV5B4PoeMVPtCuCyq",
		"2N9iShvRYeLpmQkuwwz41WCmDinkfLgg3X3",
		"2NAZSywv2d6Vztd1FcF11ZXHt5TqEmWnLHS",
	},
	"invalid": []string{
		"2NAZSywv2d6Vztd1FcF11ZXHt5TqEmWnLHs",
		"mraKKEsKHn7BA6qTcZA5ZRbXnfGXSRRFGw",
		"GEzKoZ59mhmpMzjNBWNoYKvLhFLAdHuL6P",
		"ASac6nDf6B6RWrPfWtg15q5yx2ipwGpdiK",
		"abcde",
		"",
	},
}

func TestValidateAddress(t *testing.T) {
	validator, _ := New([]string{"P2PKH", "P2SH"})
	validatorP2PKH, _ := New([]string{"P2PKH"})
	validatorP2SH, _ := New([]string{"P2SH"})
	for k, v := range cases {
		switch k {
		case "mainnetP2PKH":
			validateAddresses(t, validator, v, false, true)
			validateAddresses(t, validatorP2PKH, v, false, true)
			validateAddresses(t, validatorP2SH, v, false, false)
			validateAddresses(t, validator, v, true, false)
		case "testnetP2PKH":
			validateAddresses(t, validator, v, true, true)
			validateAddresses(t, validatorP2PKH, v, true, true)
			validateAddresses(t, validatorP2SH, v, true, false)
			validateAddresses(t, validator, v, false, false)
		case "mainnetP2SH":
			validateAddresses(t, validator, v, false, true)
			validateAddresses(t, validatorP2PKH, v, false, false)
			validateAddresses(t, validatorP2SH, v, false, true)
			validateAddresses(t, validator, v, true, false)
		case "testnetP2SH":
			validateAddresses(t, validator, v, true, true)
			validateAddresses(t, validatorP2PKH, v, true, false)
			validateAddresses(t, validatorP2SH, v, true, true)
			validateAddresses(t, validator, v, false, false)
		case "invalid":
			validateAddresses(t, validator, v, true, false)
			validateAddresses(t, validatorP2PKH, v, true, false)
			validateAddresses(t, validatorP2SH, v, true, false)
			validateAddresses(t, validator, v, false, false)
		}
	}
}

func validateAddresses(t *testing.T, validator *Validator, addresses []string, isTestnet bool, expect bool) {
	for _, a := range addresses {
		result := validator.ValidateAddress(a, isTestnet)
		assert.Equal(t, expect, result.IsValid)
	}
}
