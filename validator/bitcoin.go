package validator

// Bitcoin ...
type Bitcoin struct{}

var _ validatorBitcoinLike = (*Bitcoin)(nil)

// ValidateAddress returns validate result of bitcoin address
func (v *Bitcoin) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := normalAddrType(v, addr, network); addrType != Unknown {
		return &Result{IsValid: true, Type: addrType}
	}

	if addrType := segwitAddrType(v, addr, network); addrType != Unknown {
		return &Result{IsValid: true, Type: addrType}
	}

	return &Result{IsValid: false, Type: Unknown}
}

// AddressVersion returns bitcoin address version according to the address type and
// network type
func (v *Bitcoin) AddressVersion(addrType AddressType, network NetworkType) byte {
	if network == Mainnet {
		if addrType == P2PKH {
			return 0
		}
		return 5
	}

	if addrType == P2PKH {
		return 111
	}
	return 196
}

// AddressHrp returns hrps of bitcoin according to the network
func (v *Bitcoin) AddressHrp(network NetworkType) string {
	if network == Mainnet {
		return "bc"
	}
	return "tb"
}

// SegwitProgramLength returns segwit program length of bitcoin
func (v *Bitcoin) SegwitProgramLength(addrType AddressType) int {
	if addrType == P2WPKH {
		return 20
	} else if addrType == P2WSH {
		return 32
	}

	panic(ErrInvalidAddrType.Error())
}
