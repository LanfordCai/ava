package validator

import "strings"

// Validator ...
type Validator interface {
	ValidateAddress(addr string, network NetworkType) *Result
}

// Segwit ...
type Segwit interface {
	AddressHrp(network NetworkType) string
	SegwitProgramLength(addrType AddressType) int
}

// SegwitAddrType ...
func SegwitAddrType(v Segwit, addr string, network NetworkType) AddressType {
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

// Prefixer ...
type Prefixer interface {
	GetPrefix(network NetworkType) string
}

// IsPrefixValid ...
func IsPrefixValid(v Prefixer, addr string, network NetworkType) bool {
	return strings.HasPrefix(addr, v.GetPrefix(network))
}

// AddressWithoutPrefix ...
func AddressWithoutPrefix(v Prefixer, addr string, network NetworkType) string {
	return strings.TrimPrefix(addr, v.GetPrefix(network))
}
