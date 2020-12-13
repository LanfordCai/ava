package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZilliqaValidateAddress(t *testing.T) {
	validator := &Zilliqa{}

	var mainnetCases = map[string]*Result{
		"zil1pyvfpng2rv2p47ha67yh0edd7eyxwyv0w8hzt6": {true, Normal, ""},
		"zil144fujqgsuvv5zyr2v23m7ndyt3artkh78kyd92": {true, Normal, ""},
		"zil13gpgtu0a2shq5ck66cmx8vk8vdzv92agkcqhsa": {true, Normal, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"zib13gpgtu0a2shq5ck66cmx8vk8vdzv92agkcqhsa":                      {false, Unknown, ""},
		"zil13gegtu0a2shq5ck66cmx8vk8vdzv92agkcqhsa":                      {false, Unknown, ""},
		"zil144fujqgsuvv5zyr2v23m7ndyt3artkh78eyd92 ":                     {false, Unknown, ""},
		"ZIL1pyvfpng2rv2p47ha67yh0edd7eyxwyv0w8hzt6":                      {false, Unknown, ""},
		"zil144fujqgsuvv5zyr2v23m7ndyt3artkh78kyd92 ":                     {false, Unknown, ""},
		"zil1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3pjxtptv": {false, Unknown, ""},
		"abcde": {false, Unknown, ""},
		"":      {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
