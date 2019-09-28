package zcash

import (
	"github.com/LanfordCai/ava/internal/bitcoinlike"
)

// Validator - Method receiver, used to validate zcash transparent addresses
type Validator struct {
	bitcoinlike.Validator
}

// New - Create a new Zcash address validator
func New(types []string) (*Validator, error) {
	v := bitcoinlike.Validator{
		ChainName:           "Zcash",
		MainnetP2PKHAddrVer: []byte{28, 184},
		MainnetP2SHAddrVer:  []byte{28, 189},
		TestnetP2PKHAddrVer: []byte{29, 37},
		TestnetP2SHAddrVer:  []byte{28, 186},
		EnabledTypes:        types,
		SupportedTypes:      []string{"P2PKH", "P2SH"},
	}

	err := v.CheckTypes(types)
	if err != nil {
		return nil, err
	}

	return &Validator{v}, nil
}
