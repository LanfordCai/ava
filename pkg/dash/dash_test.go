package dash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type network int

const (
	testnet network = iota + 1
	mainnet
)

var cases = map[string][]string{
	"mainnetP2PKH": []string{
		"XpMfHazu1y8bgnP8csW73LWULGKgs4XH6U",
		"Xdk9TLNrELrM4jpu6m15zZvurG66sZKi6C",
		"XripctRDfASrofg4BKco41Zqvm1KaCdQpk",
	},
	"mainnetP2SH": []string{
		"7Z5BvydGVzgbX9xqWEb1JtF9TkToG8htcV",
	},
	"testnetP2PKH": []string{
		"yYmrsYP3XYMZr1cGtha3QzmuNB1C7CfyhV",
		"yQgGqVdasi5jGfweJ84HJz4qp4ac5G2gxG",
		"yUr7iXCNWDpF96W2kurtgszsrjf1XjDaUh",
	},
	"testnetP2SH": []string{
		"8oxdi3uiMFJ8SEob6qW76YBgg5en4352Cc",
		"92WthWvUyvzcJokeo3tyNgNZqzYmhbHM8N",
	},
	"invalid": []string{
		"yYrsYP3XYMZr1cGtha3QzmuNB1C7CfyhV",
		"7Z5BvydGVzgbX9xqWEb1JtF9TkToG8htcv",
		"Xdk9TLNreLrM4jpu6m15zZvurG66sZKi6C",
		"abcde",
		"",
	},
}

func TestValidateAddress(t *testing.T) {
	validator, _ := New([]string{"P2PKH", "P2SH"})
	validatorP2PKH, _ := New([]string{"P2PKH"})
	validatorP2SH, _ := New([]string{"P2SH"})
	for k, v := range cases {
		switch k {
		case "mainnetP2PKH":
			validateAddresses(t, validator, v, mainnet, true)
			validateAddresses(t, validatorP2PKH, v, mainnet, true)
			validateAddresses(t, validatorP2SH, v, mainnet, false)
			validateAddresses(t, validator, v, testnet, false)
		case "testnetP2PKH":
			validateAddresses(t, validator, v, testnet, true)
			validateAddresses(t, validatorP2PKH, v, testnet, true)
			validateAddresses(t, validatorP2SH, v, testnet, false)
			validateAddresses(t, validator, v, mainnet, false)
		case "mainnetP2SH":
			validateAddresses(t, validator, v, mainnet, true)
			validateAddresses(t, validatorP2PKH, v, mainnet, false)
			validateAddresses(t, validatorP2SH, v, mainnet, true)
			validateAddresses(t, validator, v, testnet, false)
		case "testnetP2SH":
			validateAddresses(t, validator, v, testnet, true)
			validateAddresses(t, validatorP2PKH, v, testnet, false)
			validateAddresses(t, validatorP2SH, v, testnet, true)
			validateAddresses(t, validator, v, mainnet, false)
		case "invalid":
			validateAddresses(t, validator, v, testnet, false)
			validateAddresses(t, validatorP2PKH, v, testnet, false)
			validateAddresses(t, validatorP2SH, v, testnet, false)
			validateAddresses(t, validator, v, mainnet, false)
		}
	}
}

func validateAddresses(t *testing.T, validator *Validator, addresses []string, network network, expect bool) {
	for _, a := range addresses {
		isValid, _ := validator.ValidateAddress(a, network == testnet)
		assert.Equal(t, expect, isValid)
	}
}
