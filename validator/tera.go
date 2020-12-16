package validator

import "strconv"

// Tera ...
type Tera struct {
	Client *TeraClient
}

var _ OnchainValidator = (*Tera)(nil)

// ValidateAddress returns validate result of stellar address
func (v *Tera) ValidateAddress(addr string, network NetworkType) *Result {
	if isValid := v.IsAddressFormatValid(addr); !isValid {
		return &Result{Success, false, Unknown, ""}
	}

	addrType, err := v.Client.GetAccount(addr)
	if err != nil {
		return &Result{Failure, false, Unknown, err.Error()}
	}

	if addrType == Unknown {
		return &Result{Success, false, Unknown, ""}
	}

	return &Result{Success, true, addrType, ""}
}

// IsAddressFormatValid ...
func (v *Tera) IsAddressFormatValid(addr string) bool {
	_, err := strconv.ParseUint(addr, 10, 32)
	return err == nil
}
