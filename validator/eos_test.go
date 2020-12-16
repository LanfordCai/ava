package validator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEOSIsAddressFormatValid(t *testing.T) {
	client := EOSClient{Endpoint: os.Getenv("AVA_EOS_ENDPOINT")}
	validator := &EOS{Client: &client}

	var validCases = []string{
		"eosnationftw",
		"atticlabeosb",
		"helloeoscnbp",
		"eosasia11111",
		"zbeosbp11111",
		"eoshuobipool",
		"big.one",
		"okcapitalbp1",
	}

	for _, addr := range validCases {
		assert.True(t, validator.IsAddrFormatValid(addr), addr)
	}

	var invalidCases = []string{
		"1YURbVuocZZZPi8LPU6GfAcKShYY7hLXbrG75v9zBXbS2zaqaHfSmGJvNEZwU3oETNZdPNxqLwR5C",
		"EOS4xXY8vZU1VGiggW5Qn7AhFcp5ti8vnWe9TVt8nUy1sUdFrLVUC",
		"eosnationft0",
		"Atticlabeosb",
		"eosasi!11111 ",
		"",
	}

	for _, addr := range invalidCases {
		assert.False(t, validator.IsAddrFormatValid(addr), addr)
	}
}
