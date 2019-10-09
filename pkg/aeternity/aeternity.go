package aeternity

import (
	"strings"

	"github.com/btcsuite/btcutil/base58"
)

const aeternityAddrPrefix = "ak_"

// Validator - Aeternity address validator
type Validator struct{}

// New - Create a Aeternity address validator
func New() *Validator {
	return &Validator{}
}

// ValidateAddress - Check an Aeternity address is valid or not
func (v *Validator) ValidateAddress(address string, isTestnet bool) (isValid bool, msg string) {
	if !strings.HasPrefix(address, aeternityAddrPrefix) {
		return false, ""
	}

	encodedPubkey := strings.TrimPrefix(address, aeternityAddrPrefix)
	decoded, _, err := base58.CheckDecode(encodedPubkey)
	if err != nil || len(decoded) != 31 {
		return false, ""
	}

	return true, ""
}
