package validator

import (
	"github.com/btcsuite/btcutil/base58"
)

// BitcoinLike ...
type BitcoinLike interface {
	Validator
	AddressVersion(addrType AddressType, network NetworkType) byte
	AddressHrp(network NetworkType) string
	SegwitProgramLength(addrType AddressType) int
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

// SegwitAddrType ...
func SegwitAddrType(v BitcoinLike, addr string, network NetworkType) AddressType {
	hrp, version, program, err := segwitAddrDecode(addr)

	if err != nil || version != 0 || hrp != v.AddressHrp(network) {
		return Unknown
	}

	if len(program) == v.SegwitProgramLength(P2WPKH) {
		return P2WPKH
	}

	if len(program) == v.SegwitProgramLength(P2WSH) {
		return P2WSH
	}

	return Unknown
}
