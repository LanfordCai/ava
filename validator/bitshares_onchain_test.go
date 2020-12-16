// +build !local

package validator

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitsharesValidateAddress(t *testing.T) {
	client := BitsharesClient{Endpoint: os.Getenv("AVA_BITSHARES_ENDPOINT")}
	validator := &Bitshares{Client: &client}

	var cases = map[string]*Result{
		"zbbts001":             {Success, true, Normal, ""},
		"beos.gateway":         {Success, true, Normal, ""},
		"huobi-bts-withdrawal": {Success, true, Normal, ""},
	}

	for addr, result := range cases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}

func TestBitsharesValidateAddress_Failure(t *testing.T) {
	client := BitsharesClient{Endpoint: "https://fakeurl"}
	validator := &Bitshares{Client: &client}

	var cases = []string{
		"zbbts001",
	}

	for _, addr := range cases {
		r := validator.ValidateAddress(addr, Mainnet)
		assert.Equal(t, Failure, r.Status)
		assert.False(t, r.IsValid)
		assert.Equal(t, Unknown, r.Type)
	}
}
