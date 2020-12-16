package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIrisValidateAddress(t *testing.T) {
	validator := &Iris{}

	var mainnetCases = map[string]*Result{
		"iaa1qjqvwsmqhm6qjgs0qsyhqc3z5f5slctp6efhhw": {Success, true, Normal, ""},
		"iaa1rpdhedqrmwtjk4lluqj6m2pvhm544h2vfu2mq8": {Success, true, Normal, ""},
		"iaa1lz9nwk6y5eeafr5c54k5dw8rwje2sj2e5pcqd6": {Success, true, Normal, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"iaa1lz9nwk6y3eeafr5c54k5dw8rwje2sj2e5pcqd6":                      {Success, false, Unknown, ""},
		"baa1lz9nwk6y5eeafr5c54k5dw8rwje2sj2e5pcqd6":                      {Success, false, Unknown, ""},
		"iaa1rpdhedqrmwtjk4lluqj6m2pvhm544h2vfu2mq8 ":                     {Success, false, Unknown, ""},
		"cosmos1e4lhppa79t40eqdgl5her60p4zqtenyj839y8z":                   {Success, false, Unknown, ""},
		"iaa1qjqvwsmqhm6qjgs0qsyhqc3z5f5slctp6efhhdew":                    {Success, false, Unknown, ""},
		"iaa1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3pjxtptv": {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
