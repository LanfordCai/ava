package validator

import "strings"

// Validator ...
type Validator interface {
	ValidateAddress(addr string, network NetworkType) *Result
}

// SS58 ...
type SS58 interface {
	Validator
	AddressType() byte
	AccountIdxLen() int
	ChecksumLen() int
	CalcChecksum(payload []byte) []byte
}

// SegwitAddress ...
type SegwitAddress interface {
	AddressHrp(network NetworkType) string
	SegwitProgramLength(addrType AddressType) int
}

// SegwitAddrType ...
func SegwitAddrType(v SegwitAddress, addr string, network NetworkType) AddressType {
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

// Bech32Address ...
type Bech32Address interface {
	AddressHrp() string
	Bech32ProgramLength() int
}

// Bech32AddrType ...
func Bech32AddrType(v Bech32Address, addr string, network NetworkType) AddressType {
	hrp, program, err := bech32AddrDecode(addr)

	if err != nil || hrp != v.AddressHrp() {
		return Unknown
	}

	if len(program) == v.Bech32ProgramLength() {
		return Normal
	}

	return Unknown
}

// CashAddress ...
type CashAddress interface {
	CashAddrType(addr string, network NetworkType) AddressType
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
