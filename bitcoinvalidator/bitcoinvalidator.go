package bitcoinvalidator

import (
	"github.com/LanfordCai/ava/bitcoinlike"
)

// BitcoinValidator - Method receiver, used to validate bitcoin addresses
type BitcoinValidator struct {
	bitcoinlike.Validator
}

// New - Create a new BitcoinValidator
func New(types []string) *BitcoinValidator {
	v := bitcoinlike.Validator{
		MainnetP2PKHAddrVer: 0,
		MainnetP2SHAddrVer:  5,
		TestnetP2PKHAddrVer: 111,
		TestnetP2SHAddrVer:  196,
		AcceptableTypes:     types,
	}

	return &BitcoinValidator{v}
}
