package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolkadotValidateAddress(t *testing.T) {
	validator := &Polkadot{}

	var validCases = map[string]*Result{
		"15j4dg5GzsL1bw2U2AWgeyAk6QTxq43V7ZPbXdAmbVLjvDCK": {Success, true, Normal, ""},
		"12xtAYsRUrmbniiWQqJtECiBQrMn8AypQcXhnQAc6RB6XkLW": {Success, true, Normal, ""},
		"14Ns6kKbCoka3MS4Hn6b7oRw9fFejG8RH5rq5j63cWUfpPDJ": {Success, true, Normal, ""},
		"16DGiP6jDwAfkAeqGfkUCtheKgUzTy7UeaiFFBAv8BwX3RhN": {Success, true, Normal, ""},
		"1vTfju3zruADh7sbBznxWCpircNp9ErzJaPQZKyrUknApRu":  {Success, true, Normal, ""},
		"1tZzPmcq8Auisttygmg9g6tPMtrh9i3b22D3tKXvde7ibRB":  {Success, true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"22xtAYsRUrmbniiWQqJtECiBQrMn8AypQcXhnQAc6RB6XkLW":     {Success, false, Unknown, ""},
		"16DGiP6jDwAfkAeqGfkUCtheKGUzTy7UeaiFFBAv8BwX3RhN":     {Success, false, Unknown, ""},
		"12xtAYsRUrmbniiWQqJtECiBQrMn8AypQcXhnQAc6RB6XeLW":     {Success, false, Unknown, ""},
		"123xtAYsRUrmbniiWQqJtECiBQrMn8AypQcXhnQAc6RB6XkLW":    {Success, false, Unknown, ""},
		"1tZzPmcq8Auisttyddddgmg9g6tPMtrh9i3b22D3tKXvde7ibRB ": {Success, false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":           {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}
}
