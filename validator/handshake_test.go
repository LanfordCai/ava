package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: testnet and P2WSH casese
func TestHandshakeValidateAddress(t *testing.T) {
	validator := &Handshake{}

	var mainnetCases = map[string]*Result{
		"hs1qlu2nssfkjt782tg2tsnrw0cus9q6we87nj32me": {true, P2WPKH, ""},
		"hs1q7z8p2jzg3sfxuttf5sjxxzlxg3rf3mte8a38aw": {true, P2WPKH, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"hs1qlu2nssfkjt782tg2tsnrw0cus9q6we87nj32m":                      {false, Unknown, ""},
		"hs2qlu2nssfkjt782tg2tsnrw0cus9q6we87nj32me":                     {false, Unknown, ""},
		"fs1qlu2nssfkjt782tg2tsnrw0cus9q6we87nj32me":                     {false, Unknown, ""},
		"tm1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3pjxtptv": {false, Unknown, ""},
		"abcde": {false, Unknown, ""},
		"":      {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
