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

	var mainnetCases = map[string]*Result{
		"zbbts001":             {Success, true, Normal, ""},
		"beos.gateway":         {Success, true, Normal, ""},
		"huobi-bts-withdrawal": {Success, true, Normal, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		".bits2020bts": {Success, false, Unknown, ""},
		"BTS4xXY8vZU1VGiggW5Qn7AhFcp5ti8vnWe9TVt8nUy1sUdFrLVUC": {Success, false, Unknown, ""},
		"binance#123": {Success, false, Unknown, ""},
		"ZBBTS001":    {Success, false, Unknown, ""},
		"":            {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
