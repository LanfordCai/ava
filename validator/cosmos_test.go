package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCosmosValidateAddress(t *testing.T) {
	validator := &Cosmos{}

	var mainnetCases = map[string]*Result{
		"cosmos1rpsdl0wdk42jjf9sp57d3eflm8q8ektuz66at3": {true, Normal, ""},
		"cosmos1e4lhppa79t40eqdgl5her60p4zqtenyj839y8z": {true, Normal, ""},
		"cosmos14dhffucsgvdtyr457022l792mp6mk6pqr96f3k": {true, Normal, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"cosmos14dhffucsgvdtyr457022l792mp6mk6pqr96f3e":                      {false, Unknown, ""},
		"cosmos34dhffucsgvdtyr457022l792mp6mk6pqr96f3k":                      {false, Unknown, ""},
		"cobmos14dhffucsgvdtyr457022l792mp6mk6pqr96f3e":                      {false, Unknown, ""},
		"cosmos14dhffucsgvdtyr457022l792mp6mk6pqr96f3k ":                     {false, Unknown, ""},
		"cosmos14dhffucsgvdtyr457022l792mp6mk6pqr96f3kab":                    {false, Unknown, ""},
		"cosmos1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3pjxtptv": {false, Unknown, ""},
		"abcde": {false, Unknown, ""},
		"":      {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
