package validator

// Dogecoin ...
type Dogecoin struct{}

var _ validatorBitcoinLike = (*Dogecoin)(nil)

// ValidateAddress returns validate result of dogecoin address
func (v *Dogecoin) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := normalAddrType(v, addr, network); addrType != Unknown {
		return &Result{IsValid: true, Type: addrType}
	}

	return &Result{IsValid: false, Type: Unknown}
}

// AddressVersion returns dogecoin address version according to the address type and
// network type
func (v *Dogecoin) AddressVersion(addrType AddressType, network NetworkType) byte {
	if network == Mainnet {
		if addrType == P2PKH {
			return 30
		}
		return 22
	}

	if addrType == P2PKH {
		return 113
	}
	return 196
}

// AddressHrp returns hrps of dogecoin according to the network
func (v *Dogecoin) AddressHrp(network NetworkType) string {
	panic(ErrUnsupported.Error())
}

// SegwitProgramLength returns segwit program length of dogecoin
func (v *Dogecoin) SegwitProgramLength(addrType AddressType) int {
	panic(ErrUnsupported.Error())
}
