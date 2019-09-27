package ripple

import (
	"github.com/LanfordCai/ava/common"
	"github.com/LanfordCai/rbase58"
)

// Validator - Ripple address validator
type Validator struct{}

// New - Create a ripple address validator
func New() *Validator {
	return &Validator{}
}

// IsValidAddress - Check a Ripple address is valid or not
func (r *Validator) IsValidAddress(address string, isTestnet bool) common.ValidationResult {
	_, version, err := rbase58.CheckDecode(address)
	if err != nil {
		return common.NewValidationResult(false, "")
	}
	return common.NewValidationResult(version == 0, "")
}
