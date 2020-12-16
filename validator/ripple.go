package validator

import (
	"github.com/LanfordCai/ava/base58check"
)

// Ripple - Ripple address validator
type Ripple struct{}

var _ Validator = (*Ripple)(nil)

// ValidateAddress - Check a Ripple address is valid or not
func (r *Ripple) ValidateAddress(addr string, network NetworkType) *Result {
	decoded, err := base58check.RippleEncoder.CheckDecode(addr)
	if err != nil || len(decoded) != 21 || decoded[0] != 0 {
		return &Result{Success, false, Unknown, ""}
	}
	return &Result{Success, true, Normal, ""}
}
