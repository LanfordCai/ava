package validator

// Bytom ...
type Bytom struct{}

var _ validatorBitcoinLike = (*Qtum)(nil)

// ValidateAddress returns validate result of bytom address
func (v *Bytom) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := segwitAddrType(v, addr, network); addrType != Unknown {
		return &Result{IsValid: true, Type: addrType}
	}
	return &Result{IsValid: false, Type: Unknown}
}

// AddressVersion returns bytom address version according to the address type and
// network type
func (v *Bytom) AddressVersion(addrType AddressType, network NetworkType) byte {
	panic(ErrUnsupported.Error())
}

// AddressHrp returns hrps of bytom according to the network
func (v *Bytom) AddressHrp(network NetworkType) string {
	if network == Mainnet {
		return "bm"
	}
	return "tm"
}

// SegwitProgramLength returns segwit program length of bytom
func (v *Bytom) SegwitProgramLength(addrType AddressType) int {
	if addrType == P2WPKH {
		return 20
	} else if addrType == P2WSH {
		return 32
	}

	panic(ErrInvalidAddrType.Error())
}
