// +build !local

package validator

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEOSValidateAddress(t *testing.T) {
	client := EOSClient{Endpoint: os.Getenv("AVA_EOS_ENDPOINT")}
	validator := &EOS{Client: &client}

	var cases = map[string]*Result{
		"eosnationftw": {Success, true, Normal, ""},
		"atticlabeosb": {Success, true, Normal, ""},
		"helloeoscnbp": {Success, true, Normal, ""},
	}

	for addr, result := range cases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}

func TestEOSIsContractDeployed(t *testing.T) {
	client := EOSClient{Endpoint: os.Getenv("AVA_EOS_ENDPOINT")}
	validator := &EOS{Client: &client}

	var cases = []string{
		"huobideposit",
	}

	for _, addr := range cases {
		deployed, err := validator.IsContractDeployed(addr)
		assert.NoError(t, err)
		assert.True(t, deployed)
	}
}

func TestEOSValidateAddress_Failure(t *testing.T) {
	client := EOSClient{Endpoint: "https://fakeurl"}
	validator := &EOS{Client: &client}

	var cases = []string{
		"eosnationftw",
	}

	for _, addr := range cases {
		r := validator.ValidateAddress(addr, Mainnet)
		assert.Equal(t, Failure, r.Status)
		assert.False(t, r.IsValid)
		assert.Equal(t, Unknown, r.Type)
	}
}
