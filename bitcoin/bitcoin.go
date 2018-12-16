package bitcoin

import (
	"github.com/btcsuite/btcutil/base58"
)

// IsValidAddress ...
// SEE: https://en.bitcoin.it/wiki/List_of_address_prefixes
func IsValidAddress(address string, isTestnet bool) bool {
	return IsP2PKH(address, isTestnet) || IsP2SH(address, isTestnet)
}

// IsP2PKH ...
func IsP2PKH(address string, isTestnet bool) bool {
	_, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}
	if isTestnet {
		return version == 111
	}
	return version == 0
}

// IsP2SH ...
func IsP2SH(address string, isTestnet bool) bool {
	_, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}
	if isTestnet {
		return version == 196
	}
	return version == 5
}
