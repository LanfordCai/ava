package ripple

import "github.com/LanfordCai/btcutil/base58"

// IsValidAddress ...
func IsValidAddress(address string) bool {
	_, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}
	return version == 0
}
