package validator

import (
	"github.com/btcsuite/btcutil/base58"
)

// BitcoinLike ...
type BitcoinLike interface {
	Validator
	AddressVersion(addrType AddressType, network NetworkType) byte
}

// NormalAddrType ...
func NormalAddrType(v BitcoinLike, addr string, network NetworkType) AddressType {
	decoded, version, err := base58.CheckDecode(addr)
	if err != nil || len(decoded) != 20 {
		return Unknown
	}

	expectP2PKH := v.AddressVersion(P2PKH, network)
	if version == expectP2PKH {
		return P2PKH
	}

	expectP2SH := v.AddressVersion(P2SH, network)
	if version == expectP2SH {
		return P2SH
	}

	return Unknown
}
