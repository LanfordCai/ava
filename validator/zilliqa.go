package validator

// Zilliqa ...
type Zilliqa struct{}

var _ Validator = (*Zilliqa)(nil)
var _ Bech32Address = (*Zilliqa)(nil)

// ValidateAddress returns validate result of iris address
func (v *Zilliqa) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := Bech32AddrType(v, addr, network); addrType != Unknown {
		return &Result{true, Normal, ""}
	}

	return &Result{false, Unknown, ""}
}

// AddressHrp returns hrp of iris according to the network
func (v *Zilliqa) AddressHrp() string {
	return "zil"
}

// Bech32ProgramLength returns program length of bytom
func (v *Zilliqa) Bech32ProgramLength() int {
	return 20
}
