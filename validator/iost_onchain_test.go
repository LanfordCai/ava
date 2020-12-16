// +build !local

package validator

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIOSTValidateAddress(t *testing.T) {
	client := IOSTClient{Endpoint: os.Getenv("AVA_IOST_ENDPOINT")}
	validator := &IOST{Client: &client}

	var mainnetCases = map[string]*Result{
		"kainiost":    {Success, true, Normal, ""},
		"_iost":       {Success, true, Normal, ""},
		"13177913280": {Success, true, Normal, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}

func TestIOSTValidateAddress_Failure(t *testing.T) {
	client := IOSTClient{Endpoint: "https://fakeurl"}
	validator := &IOST{Client: &client}

	var cases = []string{
		"huobi_iost",
	}

	for _, addr := range cases {
		r := validator.ValidateAddress(addr, Mainnet)
		assert.Equal(t, Failure, r.Status)
		assert.False(t, r.IsValid)
		assert.Equal(t, Unknown, r.Type)
	}
}
