package dogecoin

import (
	"github.com/LanfordCai/ava/internal/bitcoinlike"
)

// Validator - Method receiver, used to validate dogecoin addresses
type Validator struct {
	bitcoinlike.Validator
}

// New - Create a new Dogecoin address validator
func New(types []string) (*Validator, error) {
	v := bitcoinlike.Validator{
		ChainName:           "Dogecoin",
		MainnetP2PKHAddrVer: []byte{30},
		MainnetP2SHAddrVer:  []byte{22},
		TestnetP2PKHAddrVer: []byte{113},
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
