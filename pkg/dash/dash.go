package dash

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
		ChainName:           "Dash",
		MainnetP2PKHAddrVer: []byte{76},
		MainnetP2SHAddrVer:  []byte{16},
		TestnetP2PKHAddrVer: []byte{140},
		TestnetP2SHAddrVer:  []byte{19},
		EnabledTypes:        types,
		SupportedTypes:      []string{"P2PKH", "P2SH"},
	}

	err := v.CheckTypes(types)
	if err != nil {
		return nil, err
	}

	return &Validator{v}, nil
}
