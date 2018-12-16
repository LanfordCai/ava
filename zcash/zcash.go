package zcash

import (
	"github.com/btcsuite/btcutil/base58"
)

// IsValidAddress ...
// SEE: https://github.com/zcash/zcash/issues/2010
func IsValidAddress(address string, isTestnet bool) bool {
	return IsP2PKH(address, isTestnet) || IsP2SH(address, isTestnet)
}

// IsP2PKH ...
func IsP2PKH(address string, isTestnet bool) bool {
	decoded, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}
	if isTestnet {
		return version == 29 && decoded[0] == 37
	}
	return version == 28 && decoded[0] == 184
}

// IsP2SH ...
func IsP2SH(address string, isTestnet bool) bool {
	decoded, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}
	if isTestnet {
		return version == 28 && decoded[0] == 186
	}
	return version == 28 && decoded[0] == 189
}
