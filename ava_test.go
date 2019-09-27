package ava

import (
	"os"
	"testing"

	"github.com/LanfordCai/ava/common"
	"github.com/stretchr/testify/assert"
)

func TestNewValidator(t *testing.T) {
	os.Setenv("BITCOIN_ENABLED_ADDR_TYPES", "P2PKH")

	v, err := NewValidator("Bitcoin")
	assert.Nil(t, err)
	result := v.ValidateAddress("1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF", false)
	assert.Equal(t, true, result.IsValid)

	_, err = NewValidator("Testchain")
	assert.EqualError(t, common.ErrUnsupportedChain, err.Error())
}
