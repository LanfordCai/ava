package ava

import (
	"github.com/LanfordCai/rbase58"
)

// Ripple - Ripple address validator
type Ripple struct{}

var _ Validator = (*Ripple)(nil)

// ValidateAddress - Check a Ripple address is valid or not
func (r *Ripple) ValidateAddress(addr string, network NetworkType) *Result {
	decoded, version, err := rbase58.CheckDecode(addr)
	if err != nil || len(decoded) != 20 || version != 0 {
		return &Result{false, Unknown, ""}
	}
	return &Result{true, Normal, ""}
}
