package validator

import "encoding/base64"

// Arweave ...
type Arweave struct{}

var _ Validator = (*Arweave)(nil)

// ValidateAddress returns validate result of arweave address
func (v *Arweave) ValidateAddress(addr string, network NetworkType) *Result {
	if len(addr) != 43 {
		return &Result{Success, false, Unknown, ""}
	}

	data, err := base64.RawURLEncoding.DecodeString(addr)
	if err != nil || len(data) != 32 {
		return &Result{Success, false, Unknown, ""}
	}

	return &Result{Success, true, Normal, ""}
}
