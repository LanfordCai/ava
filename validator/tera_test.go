package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeraIsAddressFormatValid(t *testing.T) {
	validator := &Tera{}

	var validCases = []string{
		"189862",
		"224577",
		"189007",
	}

	for _, addr := range validCases {
		assert.True(t, validator.IsAddressFormatValid(addr, Mainnet), addr)
	}

	var invalidCases = []string{
		"-12345",
		"11111111111111111111111111111",
		"gDDYS6MS7EBKGC4YNSWZQZGCMQ5LKGJ5R3VAT4JUCTF6XWMNQJNYWIMF",
		"th_2SYg4CBCZpCBTnYncSzE25sRUMN4uFDDKc11pMD7u73J6xaP1g",
		"1234.01",
		"abcde",
		"",
	}

	for _, addr := range invalidCases {
		assert.False(t, validator.IsAddressFormatValid(addr, Mainnet), addr)
	}
}
