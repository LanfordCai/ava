package validator

// Handshake ...
type Handshake struct{}

var _ BitcoinLike = (*Handshake)(nil)

// ValidateAddress returns validate result of handshake address
func (v *Handshake) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := SegwitAddrType(v, addr, network); addrType != Unknown {
		return &Result{IsValid: true, Type: addrType}
	}
	return &Result{IsValid: false, Type: Unknown}
}

// AddressVersion returns handshake address version according to the address type and
// network type
func (v *Handshake) AddressVersion(addrType AddressType, network NetworkType) byte {
	panic(ErrUnsupported.Error())
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
