package aeternity

import (
	"strings"

	"github.com/LanfordCai/ava/internal/common"
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
func (v *Validator) ValidateAddress(address string, isTestnet bool) common.ValidationResult {
	if !strings.HasPrefix(address, aeternityAddrPrefix) {
		return common.NewValidationResult(false, "")
	}

	encodedPubkey := strings.TrimPrefix(address, aeternityAddrPrefix)
	decoded, _, err := base58.CheckDecode(encodedPubkey)
	if err != nil || len(decoded) != 31 {
		return common.NewValidationResult(false, "")
	}

	return common.NewValidationResult(true, "")
}
