package bytom

import (
	"github.com/LanfordCai/ava/bitcoinlike"
)

// Validator - Method receiver, used to validate bitcoin addresses
type Validator struct {
	bitcoinlike.Validator
}

// New - Create a new bitcoin address validator
func New(types []string) (*Validator, error) {
	v := bitcoinlike.Validator{
		ChainName:        "Bytom",
		MainnetBech32HRP: "bm",
		TestnetBech32HRP: "tm",
		EnabledTypes:     types,
		SupportedTypes:   []string{"P2WPKH", "P2WSH"},
	}

	err := v.CheckTypes(types)
	if err != nil {
		return nil, err
	}

	return &Validator{v}, nil
}
