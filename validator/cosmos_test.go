package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCosmosValidateAddress(t *testing.T) {
	validator := &Cosmos{}

	var mainnetCases = map[string]*Result{
		"cosmos1rpsdl0wdk42jjf9sp57d3eflm8q8ektuz66at3": {Success, true, Normal, ""},
		"cosmos1e4lhppa79t40eqdgl5her60p4zqtenyj839y8z": {Success, true, Normal, ""},
		"cosmos14dhffucsgvdtyr457022l792mp6mk6pqr96f3k": {Success, true, Normal, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"cosmos14dhffucsgvdtyr457022l792mp6mk6pqr96f3e":                      {Success, false, Unknown, ""},
		"cosmos34dhffucsgvdtyr457022l792mp6mk6pqr96f3k":                      {Success, false, Unknown, ""},
		"cobmos14dhffucsgvdtyr457022l792mp6mk6pqr96f3e":                      {Success, false, Unknown, ""},
		"cosmos14dhffucsgvdtyr457022l792mp6mk6pqr96f3k ":                     {Success, false, Unknown, ""},
		"cosmos14dhffucsgvdtyr457022l792mp6mk6pqr96f3kab":                    {Success, false, Unknown, ""},
		"cosmos1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3pjxtptv": {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
