// +build !local

package validator

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOasisValidateAddress(t *testing.T) {
	client := RosettaClient{Endpoint: os.Getenv("AVA_OASIS_ENDPOINT")}
	validator := &Oasis{Client: &client}

	var cases = map[string]*Result{
		"oasis1qqs6ylpfurhf6gc9mw232fkmrt3d0673lyzc5xf2": {Success, true, Normal, ""},
		"oasis1qrad7s7nqm4gvyzr8yt2rdk0ref489rn3vn400d6": {Success, true, Normal, ""},
		"oasis1qp29h8ykmxet46eqzw0wennrmmy4al3xzv37m3ca": {Success, true, Normal, ""},
		"oasis1qp29h8ykmxet46eqzw0wennrmmy4al3xzv37m3ce": {Success, false, Unknown, ""},
	}

	for addr, result := range cases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}

func TestOasisValidateAddress_Failure(t *testing.T) {
	client := RosettaClient{Endpoint: "https://fakeurl"}
	validator := &Oasis{Client: &client}

	var cases = []string{
		"oasis1qrad7s7nqm4gvyzr8yt2rdk0ref489rn3vn400d6",
	}

	for _, addr := range cases {
		r := validator.ValidateAddress(addr, Mainnet)
		assert.Equal(t, Failure, r.Status)
		assert.False(t, r.IsValid)
		assert.Equal(t, Unknown, r.Type)
	}
}
