package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeepinValidateAddress(t *testing.T) {
	validator := &Zeepin{}

	var mainnetCases = map[string]*Result{
		"ZQfyK6J8sRtKT5XECyrPDV4KsAXPWfYARd": {Success, true, Normal, ""},
		"ZJSmkuhUL6GUaVwzapCaXyYPT2cBYgMQ4n": {Success, true, Normal, ""},
		"ZD8HgfakgSXjk9XN2c5SS5qbEZp94TEmE6": {Success, true, Normal, ""},
		"ZHeWYS7Kc2Nti6gT8KKec6a4ohGNNcTxVF": {Success, true, Normal, ""},
		"ZakH4H7PG3uXMj1EFFmaz3s1ULYvgTuSRU": {Success, true, Normal, ""},
		"ZEJQUH8MhVRcNY8o4zECS8ZGyqDJYEBTGJ": {Success, true, Normal, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"ZWC6BzJXCGwZ6ePv71isGgFjVoQuHME7PZ5":                                    {Success, false, Unknown, ""},
		"ZWC6BzJXGGwZ6ePv76sGgFjVoQuHME7PZ4":                                     {Success, false, Unknown, ""},
		"0100000000000000000000000000000000000000":                               {Success, false, Unknown, ""},
		"AVee5uM9eHNSfcm1nVrS6SxCe6gKMtc6hJ":                                     {Success, false, Unknown, ""},
		"NBVH-33D3BQ2TWUR6TJNZTZOVR2M3A-ZDUZ-RDRK-R4EA":                          {Success, false, Unknown, ""},
		"ATscE4ZCHfouSJ7mXacbYQAsSFZgUHPw22FATscE4ZCHfouSJ7mXacbYQAsSFZgUHPw22F": {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}
}
