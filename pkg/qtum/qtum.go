package qtum

import (
	"github.com/LanfordCai/ava/internal/bitcoinlike"
)

// Validator - Method receiver, used to validate qtum addresses
type Validator struct {
	bitcoinlike.Validator
}

// New - Create a new Qtum address validator
func New(types []string) (*Validator, error) {
	v := bitcoinlike.Validator{
		ChainName:           "Qtum",
		MainnetP2PKHAddrVer: []byte{58},
		MainnetP2SHAddrVer:  []byte{50},
		TestnetP2PKHAddrVer: []byte{120},
		TestnetP2SHAddrVer:  []byte{110},
		EnabledTypes:        types,
		SupportedTypes:      []string{"P2PKH", "P2SH"},
	}

	err := v.CheckTypes(types)
	if err != nil {
		return nil, err
	}

	return &Validator{v}, nil
}
