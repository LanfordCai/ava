package neo

import (
	"github.com/btcsuite/btcutil/base58"
)

// Validator - Neo address validator
type Validator struct{}

const addrVer = 23

// New - Create a Neo address validator
func New() *Validator {
	return &Validator{}
}

// ValidateAddress - Check a Neo address is valid or not
func (n *Validator) ValidateAddress(address string, isTestnet bool) (isValid bool, msg string) {
	_, version, err := base58.CheckDecode(address)
	if err != nil {
		return false, ""
	}
	return version == addrVer, ""
}
