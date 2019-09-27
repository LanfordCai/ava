package litecoin

import (
	"github.com/LanfordCai/ava/bitcoinlike"
)

// Validator - Method receiver, used to validate litecoin addresses
type Validator struct {
	bitcoinlike.Validator
}

// New - Create a new LitecoinValidator
func New(types []string) (*Validator, error) {
	v := bitcoinlike.Validator{
		ChainName:           "Litecoin",
		MainnetP2PKHAddrVer: []byte{48},
		MainnetP2SHAddrVer:  []byte{50},
		TestnetP2PKHAddrVer: []byte{111},
		TestnetP2SHAddrVer:  []byte{58},
		EnabledTypes:        types,
		SupportedTypes:      []string{"P2PKH", "P2SH"},
	}

	err := v.CheckTypes(types)
	if err != nil {
		return nil, err
	}

	return &Validator{v}, nil
}
