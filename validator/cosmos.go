package validator

// Cosmos ...
type Cosmos struct{}

var _ Validator = (*Cosmos)(nil)
var _ Bech32Address = (*Cosmos)(nil)

// ValidateAddress returns validate result of cosmos address
func (v *Cosmos) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := Bech32AddrType(v, addr, network); addrType != Unknown {
		return &Result{true, Normal, ""}
	}

	return &Result{false, Unknown, ""}
}

// AddressHrp returns hrp of cosmos according to the network
func (v *Cosmos) AddressHrp() string {
	return "cosmos"
}

// Bech32ProgramLength returns program length of bytom
func (v *Cosmos) Bech32ProgramLength() int {
	return 20
}
