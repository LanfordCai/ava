package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitsharesIsAddressFormatValid(t *testing.T) {
	validator := &Bitshares{}

	var validCases = []string{
		"zbbts001",
		"beos.gateway",
		"huobi-bts-withdrawal",
	}

	for _, addr := range validCases {
		assert.True(t, validator.IsAddressFormatValid(addr), addr)
	}

	var invalidCases = []string{
		".bits2020bts",
		"BTS4xXY8vZU1VGiggW5Qn7AhFcp5ti8vnWe9TVt8nUy1sUdFrLVUC",
		"binance#123",
		"ZBBTS001",
		"",
	}

	for _, addr := range invalidCases {
		assert.False(t, validator.IsAddressFormatValid(addr), addr)
	}
}
