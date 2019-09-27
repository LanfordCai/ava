package sia

import (
	"encoding/hex"

	"github.com/LanfordCai/ava/common"
	"golang.org/x/crypto/blake2b"
)

// Validator - Sia address validator
type Validator struct{}

// New - Create a Sia address validator
func New() *Validator {
	return &Validator{}
}

// ValidateAddress - Check a Sia address is valid or not
func (s *Validator) ValidateAddress(address string, isTestnet bool) common.ValidationResult {
	unlockhashWithChecksum, err := hex.DecodeString(address)
	if err != nil || len(unlockhashWithChecksum) != 38 {
		return common.NewValidationResult(false, "")
	}
	unlockhash := unlockhashWithChecksum[0:32]
	checksum256 := blake2b.Sum256(unlockhash)

	var validChecksum [6]byte
	copy(validChecksum[:], checksum256[:6])

	var checksum [6]byte
	copy(checksum[:], unlockhashWithChecksum[32:])
	return common.NewValidationResult(checksum == validChecksum, "")
}
