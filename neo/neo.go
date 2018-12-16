package neo

import "github.com/btcsuite/btcutil/base58"

// IsValidAddress ...
func IsValidAddress(address string) bool {
	_, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}
	return version == 23
}
