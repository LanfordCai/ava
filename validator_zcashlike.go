package ava

import (
	"bytes"

	"github.com/btcsuite/btcutil/base58"
)

// ValidatorZcashLike ...
type ValidatorZcashLike interface {
	Validator
	AddressVersion(addrType AddressType, network NetworkType) []byte
}

func zcashlikeNormalAddrType(v ValidatorZcashLike, addr string, network NetworkType) AddressType {
	decoded, version, err := base58.CheckDecode(addr)
	if err != nil {
		return Unknown
	}

	if len(decoded) != 21 {
		return Unknown
	}

	expectP2PKH := v.AddressVersion(P2PKH, network)
	if bytes.Compare([]byte{version, decoded[0]}, expectP2PKH) == 0 {
		return P2PKH
	}

	expectP2SH := v.AddressVersion(P2SH, network)
	if bytes.Compare([]byte{version, decoded[0]}, expectP2SH) == 0 {
		return P2SH
	}

	return Unknown
}
