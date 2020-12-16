package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTronValidateAddress(t *testing.T) {
	validator := &Tron{}

	var validCases = map[string]*Result{
		"TKTcfBEKpp5ZRPwmiZ8SfLx8W7CDZ7PHCY": {Success, true, Normal, ""},
		"TW73UWp1a5s8Zs36D6LkQuBp27dz2VRg1a": {Success, true, Normal, ""},
		"TLyqzVGLV1srkB7dToTAEqgDSfPtXRJZYH": {Success, true, Normal, ""},
		"TCFLL5dx5ZJdKnWuesXxi1VPwjLVmWZZy9": {Success, true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"TKTcfBEKpp5ZRPwmiZ8SfLx8W7CDZ7PHCY ":        {Success, false, Unknown, ""},
		"AN2aGVjEBi12hLWg65iGpraTn45h4jcLdU":         {Success, false, Unknown, ""},
		"6way2gX9m8hjhWFSSaemoZk1tLDFLidvH":          {Success, false, Unknown, ""},
		"412b71baefd4359f683c70e8ed81f5fb9aeff2cd89": {Success, false, Unknown, ""},
		"invalid": {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}
}
