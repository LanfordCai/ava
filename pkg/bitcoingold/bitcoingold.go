package bitcoingold

import (
	"github.com/LanfordCai/ava/internal/bitcoinlike"
)

// Validator - Method receiver, used to validate bitcoin addresses
type Validator struct {
	bitcoinlike.Validator
}

// New - Create a new bitcoin address validator
func New(types []string) (*Validator, error) {
	v := bitcoinlike.Validator{
		ChainName:           "BitcoinGold",
		MainnetP2PKHAddrVer: []byte{38},
		MainnetP2SHAddrVer:  []byte{23},
		TestnetP2PKHAddrVer: []byte{111},
		TestnetP2SHAddrVer:  []byte{196},
		EnabledTypes:        types,
		SupportedTypes:      []string{"P2PKH", "P2SH"},
	}

	err := v.CheckTypes(types)
	if err != nil {
		return nil, err
	}

	return &Validator{v}, nil
}
