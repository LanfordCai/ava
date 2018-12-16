package sia

import (
	"encoding/hex"

	"golang.org/x/crypto/blake2b"
)

// IsValidAddress ...
func IsValidAddress(address string) bool {
	unlockhashWithChecksum, err := hex.DecodeString(address)
	if err != nil || len(unlockhashWithChecksum) != 38 {
		return false
	}
	unlockhash := unlockhashWithChecksum[0:32]
	checksum256 := blake2b.Sum256(unlockhash)

	var validChecksum [6]byte
	copy(validChecksum[:], checksum256[:6])

	var checksum [6]byte
	copy(checksum[:], unlockhashWithChecksum[32:])
	return checksum == validChecksum
}
