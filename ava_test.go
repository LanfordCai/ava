package ava

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidatorWithEnabledTypes(t *testing.T) {
	os.Setenv("AVA_BITCOIN_ENABLED_ADDR_TYPES", "P2PKH")

	v, err := NewValidator("Bitcoin")
	assert.Nil(t, err)
	isValid, _ := v.ValidateAddress("1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF", false)
	assert.Equal(t, true, isValid)

	_, err = NewValidator("Testchain")
	assert.EqualError(t, ErrUnsupportedChain, err.Error())
}

func TestNewValidatorWithContractWhitelist(t *testing.T) {
	os.Setenv("AVA_EOS_CONTRACT_WHITELIST", "huobideposit")

	v, err := NewValidator("EOS")
	assert.Nil(t, err)
	isValid, _ := v.ValidateAddress("huobideposit", false)
	assert.Equal(t, true, isValid)

	isValid, _ = v.ValidateAddress("pxneosincome", false)
	assert.Equal(t, true, isValid)
}
