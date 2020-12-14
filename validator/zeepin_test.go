package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeepinValidateAddress(t *testing.T) {
	validator := &Zeepin{}

	var mainnetCases = map[string]*Result{
		"ZQfyK6J8sRtKT5XECyrPDV4KsAXPWfYARd": {true, Normal, ""},
		"ZJSmkuhUL6GUaVwzapCaXyYPT2cBYgMQ4n": {true, Normal, ""},
		"ZD8HgfakgSXjk9XN2c5SS5qbEZp94TEmE6": {true, Normal, ""},
		"ZHeWYS7Kc2Nti6gT8KKec6a4ohGNNcTxVF": {true, Normal, ""},
		"ZakH4H7PG3uXMj1EFFmaz3s1ULYvgTuSRU": {true, Normal, ""},
		"ZEJQUH8MhVRcNY8o4zECS8ZGyqDJYEBTGJ": {true, Normal, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"ZWC6BzJXCGwZ6ePv71isGgFjVoQuHME7PZ5":                                    {false, Unknown, ""},
		"ZWC6BzJXGGwZ6ePv76sGgFjVoQuHME7PZ4":                                     {false, Unknown, ""},
		"0100000000000000000000000000000000000000":                               {false, Unknown, ""},
		"AVee5uM9eHNSfcm1nVrS6SxCe6gKMtc6hJ":                                     {false, Unknown, ""},
		"NBVH-33D3BQ2TWUR6TJNZTZOVR2M3A-ZDUZ-RDRK-R4EA":                          {false, Unknown, ""},
		"ATscE4ZCHfouSJ7mXacbYQAsSFZgUHPw22FATscE4ZCHfouSJ7mXacbYQAsSFZgUHPw22F": {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}
}
