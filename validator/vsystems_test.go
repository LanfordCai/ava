package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NOTE: no testnet cases
func TestVsystemsValidateAddress(t *testing.T) {
	validator := &Vsystems{}

	var mainnetCases = map[string]*Result{
		"ARF12jvtjz9caUFmiwBeRe1SPRGQhUWKrtd": {true, Normal, ""},
		"ARBQTCYws5FZAVtA1ZFLsGhBtPymr4Hp5CX": {true, Normal, ""},
		"AR45wyKHZnmt7ujqJRT7b4hSk9wX1bjwDkz": {true, Normal, ""},
		"ARCVpcq2i6rQ7kkzNeJ1jsMec6TLmC7RNHn": {true, Normal, ""},
		"AR95ZdHV3Wx759r5io7gmb2GUz9vRtFF1F8": {true, Normal, ""},
		"AREExiJHmLb15ePMTyajnt4wb2bD4BENsM4": {true, Normal, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"AUEef5yocCa7DND878N2FTAH35QKACAtqbH": {true, Normal, ""},
		"AU3JbcU4nQfDUn5h9uhqzmpXNibP72Lhomh": {true, Normal, ""},
		"AU668Z1DLWqsWNxhdv4ay4KNjxEyPGtqHqy": {true, Normal, ""},
		"AUCDnRnuQ6ZLeakgPdESeiThj87Nw9SgLcF": {true, Normal, ""},
		"AUFLuQT2ynr1NgarK6SunEorEszLL2U97Nd": {true, Normal, ""},
		"ATvEDmb1KwSvuWCrz7epdTm5UD4xyc31qsz": {true, Normal, ""},
	}

	for addr, result := range testnetCases {
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"ATscE4ZCHfouSJ7mXacbYQAsSFZgUHPw22E":                                    {false, Unknown, ""},
		"BTxHp2AhVkB6RxafH7nAWrAAaB5qgPAVd24":                                    {false, Unknown, ""},
		"TBVH33D3BQ2TWUR6TJNZTZOVR2M3AZDUZRDRKR4E":                               {false, Unknown, ""},
		"NBVH-33D3BQ2TWUR6TJNZTZOVR2M3A-ZDUZ-RDRK-R4EA":                          {false, Unknown, ""},
		"ATscE4ZCHfouSJ7mXacbYQAsSFZgUHPw22FATscE4ZCHfouSJ7mXacbYQAsSFZgUHPw22F": {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}
}
