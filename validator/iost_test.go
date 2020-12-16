package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIOSTIsAddressFormatValid(t *testing.T) {
	validator := &IOST{}

	var validCases = []string{
		"kainiost",
		"ukay66",
		"huobi_iost",
		"13177913280",
		"_iost",
		"y2a_yfec2st",
	}

	for _, addr := range validCases {
		assert.True(t, validator.IsAddressFormatValid(addr), addr)
	}

	var invalidCases = []string{
		"1YURbVuocZZZPi8LPU6GfAcKShYY7hLXbrG75v9zBXbS2zaqaHfSmGJvNEZwU3oETNZdPNxqLwR5C",
		"IOST4xXY8vZU1VGiggW5Qn7AhFcp5ti8vnWe9TVt8nUy1sUdFrLVUC",
		"big.one",
		"huoBI_IOST",
		"0001 hello",
		"",
	}

	for _, addr := range invalidCases {
		assert.False(t, validator.IsAddressFormatValid(addr), addr)
	}
}
