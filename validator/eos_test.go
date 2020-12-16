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
	validator := EOS{Client: &client}

	var mainnetCases = map[string]*Result{
		"eosnationftw": {Success, true, Normal, ""},
		"atticlabeosb": {Success, true, Normal, ""},
		"helloeoscnbp": {Success, true, Normal, ""},
		"eosasia11111": {Success, true, Normal, ""},
		"zbeosbp11111": {Success, true, Normal, ""},
		"eoshuobipool": {Success, true, Normal, ""},
		"big.one":      {Success, true, Normal, ""},
		"okcapitalbp1": {Success, true, Normal, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"1YURbVuocZZZPi8LPU6GfAcKShYY7hLXbrG75v9zBXbS2zaqaHfSmGJvNEZwU3oETNZdPNxqLwR5C": {Success, false, Unknown, ""},
		"EOS4xXY8vZU1VGiggW5Qn7AhFcp5ti8vnWe9TVt8nUy1sUdFrLVUC":                         {Success, false, Unknown, ""},
		"eosnationft0":  {Success, false, Unknown, ""},
		"Atticlabeosb":  {Success, false, Unknown, ""},
		"eosasi!11111 ": {Success, false, Unknown, ""},
		"":              {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
