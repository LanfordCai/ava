package validator

// Handshake ...
type Handshake struct{}

var _ Validator = (*Handshake)(nil)
var _ SegwitAddress = (*Handshake)(nil)

// ValidateAddress returns validate result of handshake address
func (v *Handshake) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := SegwitAddrType(v, addr, network); addrType != Unknown {
		return &Result{IsValid: true, Type: addrType}
	}
	return &Result{IsValid: false, Type: Unknown}
}

// AddressHrp returns hrps of handshake according to the network
func (v *Handshake) AddressHrp(network NetworkType) string {
	if network == Mainnet {
		return "hs"
	}
	return "ts"
}

// SegwitProgramLength returns segwit program length of handshake
func (v *Handshake) SegwitProgramLength(addrType AddressType) int {
	if addrType == P2WPKH {
		return 20
	} else if addrType == P2WSH {
		return 32
	}

	panic(ErrInvalidAddrType.Error())
}
