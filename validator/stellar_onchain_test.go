// +build !local

package validator

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStellarValidateAddress(t *testing.T) {
	client := StellarClient{Endpoint: os.Getenv("AVA_STELLAR_ENDPOINT")}
	validator := &Stellar{Client: &client}

	var validCases = map[string]*Result{
		"GAI3GJ2Q3B35AOZJ36C4ANE3HSS4NK7WI6DNO4ZSHRAX6NG7BMX6VJER": {Success, true, Normal, ""},
		"GBLKRATZODTSJNU7XTB5HY5VAAN63CPRT77UYZT2VLCNXE7F3YHSW22M": {Success, true, Normal, ""},
		"GDDYS6MS7EBKGC4YNSWZQZGCMQ5LKGJ5R3VAT4JUCTF6XWMNQJNYWIMF": {Success, true, Normal, ""},
		"GBWEQO3QPKKC7W6UMK7BWNBHBLIL54MNI5MSOMRB3TURYWJQ64XD3EN3": {Success, true, Normal, ""},
		"GALAXY2447FJMEEQ2RIH7S42QAWVFQVE3DPOEYLTUNRNAEENBM6IO6EH": {Success, true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"GBWEQO3QPKKC7W6UMK7BWNBHBLIL54MNI5MSOMRB3URYWJQ64XD3EN3":  {Success, false, Unknown, ""},
		"GALAXY2447FJMEEQ2RIH7S42QAWVFQVE3DPOEYLTUNRNAEENBM6IO6EF": {Success, false, Unknown, ""},
		"gDDYS6MS7EBKGC4YNSWZQZGCMQ5LKGJ5R3VAT4JUCTF6XWMNQJNYWIMF": {Success, false, Unknown, ""},
		"GDDYS6MS7EBKGC4Y5LKGJ5R3VAT4JUCTF6XWMNQJNYWIMF":           {Success, false, Unknown, ""},
		"th_2SYg4CBCZpCBTnYncSzE25sRUMN4uFDDKc11pMD7u73J6xaP1g ":   {Success, false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":               {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
