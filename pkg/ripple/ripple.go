package ripple

import (
	"github.com/LanfordCai/rbase58"
)

// Validator - Ripple address validator
type Validator struct{}

// New - Create a ripple address validator
func New() *Validator {
	return &Validator{}
}

// ValidateAddress - Check a Ripple address is valid or not
func (r *Validator) ValidateAddress(address string, isTestnet bool) (isValid bool, msg string) {
	_, version, err := rbase58.CheckDecode(address)
	if err != nil {
		return false, ""
	}
	return version == 0, ""
}
