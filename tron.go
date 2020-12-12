package ava

import "github.com/btcsuite/btcutil/base58"

// Tron - Tron address validator
type Tron struct{}

var _ Validator = (*Tron)(nil)

// ValidateAddress - Check a Ripple address is valid or not
func (r *Tron) ValidateAddress(addr string, network NetworkType) *Result {
	decoded, version, err := base58.CheckDecode(addr)
	if err != nil || len(decoded) != 20 || version != 65 {
		return &Result{false, Unknown, ""}
	}
	return &Result{true, Normal, ""}
}
