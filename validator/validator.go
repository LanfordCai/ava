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

// OnchainValidator ...
type OnchainValidator interface {
	Validator
	IsAddressFormatValid(addr string, network NetworkType) bool
}

// ContractChecker ...
type ContractChecker interface {
	IsContractDeployed(addr string) (bool, error)
}

// RosettaValidator ...
type RosettaValidator interface {
	OnchainValidator
	NetworkIdentifier(network NetworkType) RosettaNetworkIdentifier
}

// RosettaNetworkIdentifier is the simplified network identifier of Rosetta
type RosettaNetworkIdentifier struct {
	Blockchain string `json:"blockchain"`
	Network    string `json:"network"`
}

// RosettaAccountIdentifier is the simplified account identifier of Rosetta
type RosettaAccountIdentifier struct {
	Address string `json:"address"`
}
