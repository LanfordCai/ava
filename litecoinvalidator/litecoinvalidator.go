package litecoinvalidator

import (
	"github.com/LanfordCai/ava/bitcoinlike"
)

// LitecoinValidator - Method receiver, used to validate litecoin addresses
type LitecoinValidator struct {
	bitcoinlike.Validator
}

// New - Create a new LitecoinValidator
func New(types []string) *LitecoinValidator {
	v := bitcoinlike.Validator{
		MainnetP2PKHAddrVer: 48,
		MainnetP2SHAddrVer:  50,
		TestnetP2PKHAddrVer: 111,
		TestnetP2SHAddrVer:  58,
		AcceptableTypes:     types,
	}

	return &LitecoinValidator{v}
}
