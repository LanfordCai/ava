package ava

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTronValidateAddress(t *testing.T) {
	validator := &Tron{}

	var validCases = map[string]*Result{
		"TKTcfBEKpp5ZRPwmiZ8SfLx8W7CDZ7PHCY": {true, Normal, ""},
		"TW73UWp1a5s8Zs36D6LkQuBp27dz2VRg1a": {true, Normal, ""},
		"TLyqzVGLV1srkB7dToTAEqgDSfPtXRJZYH": {true, Normal, ""},
		"TCFLL5dx5ZJdKnWuesXxi1VPwjLVmWZZy9": {true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"TKTcfBEKpp5ZRPwmiZ8SfLx8W7CDZ7PHCY ":        {false, Unknown, ""},
		"AN2aGVjEBi12hLWg65iGpraTn45h4jcLdU":         {false, Unknown, ""},
		"6way2gX9m8hjhWFSSaemoZk1tLDFLidvH":          {false, Unknown, ""},
		"412b71baefd4359f683c70e8ed81f5fb9aeff2cd89": {false, Unknown, ""},
		"invalid": {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}
}
