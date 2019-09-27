package ripple

import "github.com/LanfordCai/btcutil/base58"

// Validator - Ripple address validator
type Validator struct{}

// New - Create a ripple address validator
func New() *Validator {
	return &Validator{}
}

// IsValidAddress - Check a Ripple address is valid or not
func (r *Validator) IsValidAddress(address string, isTestnet bool) bool {
	_, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}
	return version == 0
}
