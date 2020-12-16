package validator

// Iris ...
type Iris struct{}

var _ Validator = (*Iris)(nil)
var _ Bech32Address = (*Iris)(nil)

// ValidateAddress returns validate result of iris address
func (v *Iris) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := Bech32AddrType(v, addr, network); addrType != Unknown {
		return &Result{Success, true, Normal, ""}
	}

	return &Result{Success, false, Unknown, ""}
}

// AddressHrp returns hrp of iris according to the network
func (v *Iris) AddressHrp() string {
	return "iaa"
}

// Bech32ProgramLength returns program length of bytom
func (v *Iris) Bech32ProgramLength() int {
	return 20
}
