package validator

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/btcsuite/btcutil/bech32"
)

type validatorBitcoinLike interface {
	validator
	AddressVersion(addrType AddressType, network NetworkType) byte
	AddressHrp(network NetworkType) string
	SegwitProgramLength(addrType AddressType) int
}

func normalAddrType(v validatorBitcoinLike, addr string, network NetworkType) AddressType {
	decoded, version, err := base58.CheckDecode(addr)
	if err != nil {
		return Unknown
	}

	if len(decoded) != 20 {
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

func segwitAddrType(v validatorBitcoinLike, addr string, network NetworkType) AddressType {
	hrp, version, program, err := segwitAddrDecode(addr)
	if err != nil {
		return Unknown
	}

	if version != 0 || hrp != v.AddressHrp(network) {
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

func segwitAddrDecode(address string) (string, byte, []byte, error) {
	hrp, data, err := bech32.Decode(address)
	if err != nil || len(data) < 1 {
		return "", 255, nil, err
	}
	version := data[0]
	program, err := bech32.ConvertBits(data[1:], 5, 8, false)
	if err != nil {
		return "", 255, nil, err
	}
	return hrp, version, program, nil
}
