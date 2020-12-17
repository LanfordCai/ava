package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKusamaValidateAddress(t *testing.T) {
	validator := &Kusama{}

	var validCases = map[string]*Result{
		"GZTZiL1F1TRYMcT7mmkUiujR6R2oVcw7bsWpWawwQJuRzdD": {Success, true, Normal, ""},
		"HEVQFCv5HWcwZhzJEX93gQfXikJQ9VgfHoQ9RvXk3KYHN29": {Success, true, Normal, ""},
		"Ek5FeXtZWxvcS5snsFMdwHZdRnaYvNYRJdBTZMiSur5bde6": {Success, true, Normal, ""},
		"DYMQA7uvnxayYxcSQywPN7jufwZovBHL48hp747me94aZdG": {Success, true, Normal, ""},
		"Ho6PiyJXo5paaZJADX3eCQUsonyvjdMaU4DJxop1CTHEPFN": {Success, true, Normal, ""},
		"FfkbFfHRvxsp4Cpr4KMhC4hXEqhd3QrnoeUsWvfsSpTyxnr": {Success, true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"FfkbFfHRvxsp4Cpr4KMhC4hXEqhd3QrnoeUsWvfsSpTyxnG":   {Success, false, Unknown, ""},
		"16DGiP6jDwAfkAeqGfkUCtheKGUzTy7UeaiFFBAv8BwX3RhN":  {Success, false, Unknown, ""},
		"Ho6PiyJXo5paaZJADX3eCQUsonyvjdMaU4DJxop1C#HEPFN":   {Success, false, Unknown, ""},
		"Ho6PiyJXo5paaZJADX3eCQUsssonyvjdMaU4DJxop1CTHEPFN": {Success, false, Unknown, ""},
		"Ho6PiyJXo5paaZJADX3eCQUsonyvjdMaU4DJxop1CTHEPFN ":  {Success, false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":        {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}
}
