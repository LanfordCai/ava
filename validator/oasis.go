package validator

import (
	"strings"

	"github.com/btcsuite/btcutil/bech32"
)

// Oasis ...
type Oasis struct {
	Client *RosettaClient
}

var _ RosettaValidator = (*Oasis)(nil)
var _ Bech32Address = (*Oasis)(nil)

// ValidateAddress returns validate result of stellar address
func (v *Oasis) ValidateAddress(addr string, network NetworkType) *Result {
	if isValid := v.IsAddressFormatValid(addr, network); !isValid {
		return &Result{Success, false, Unknown, ""}
	}

	addrType, err := v.Client.GetAccount(v.NetworkIdentifier(network), addr)
	if err != nil {
		return &Result{Failure, false, Unknown, err.Error()}
	}

	if addrType == Unknown {
		return &Result{Success, false, Unknown, ""}
	}

	return &Result{Success, true, addrType, ""}
}

// IsAddressFormatValid ...
func (v *Oasis) IsAddressFormatValid(addr string, network NetworkType) bool {
	if strings.ToLower(addr) != addr {
		return false
	}

	hrp, data, err := bech32.Decode(addr)
	if err != nil || hrp != v.AddressHrp() || len(data) != v.Bech32ProgramLength() {
		return false
	}

	return true
}

// AddressHrp ...
func (v *Oasis) AddressHrp() string {
	return "oasis"
}

// Bech32ProgramLength ...
func (v *Oasis) Bech32ProgramLength() int {
	return 34
}

// NetworkIdentifier ...
func (v *Oasis) NetworkIdentifier(network NetworkType) RosettaNetworkIdentifier {
	if network == Mainnet {
		return RosettaNetworkIdentifier{
			Blockchain: "Oasis",
			Network:    "a4dc2c4537992d6d2908c9779927ccfee105830250d903fd1abdfaf42cb45631",
		}
	}

	return RosettaNetworkIdentifier{
		Blockchain: "Oasis",
		Network:    "086a764a7a748eb6a2a3b046f152caf7e1cc9713478ce0565df253e1c5872963",
	}
}
