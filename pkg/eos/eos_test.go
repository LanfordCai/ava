package eos

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = map[string][]string{
	"invalidFormat": []string{
		"007",
		"abc-abc",
		"gogo&11",
		"pxneosincome1",
		"hUOBIDEPOSIT",
	},
	"unregistered": []string{
		"huo1ideposie",
		"pxneosincomd",
	},
	"normal": []string{
		"pxneosincome",
	},
	"contract": []string{
		"huobideposit",
	},
}

func TestValidateAddress(t *testing.T) {
	baseURL := os.Getenv("AVA_EOS_BASE_URL")
	validator := New(baseURL, nil)
	validatorWithWhitelist := New(baseURL, []string{"huobideposit"})

	for k, v := range cases {
		switch k {
		case "invalidFormat", "unregistered":
			validateAddress(t, validator, v, false)
		case "normal":
			validateAddress(t, validator, v, true)
		case "contract":
			validateAddress(t, validator, v, false)
			validateAddress(t, validatorWithWhitelist, v, true)
		}
	}
}

func validateAddress(t *testing.T, validator *Validator, addresses []string, expect bool) {
	t.Helper()
	for _, a := range addresses {
		isValid, _ := validator.ValidateAddress(a, false)
		assert.Equal(t, expect, isValid)
	}
}
