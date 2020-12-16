package validator

import (
	"regexp"
)

// IOST address validator
type IOST struct {
	Client *IOSTClient
}

var _ OnchainValidator = (*IOST)(nil)

// ValidateAddress ...
func (e *IOST) ValidateAddress(addr string, network NetworkType) *Result {
	if isValid := e.IsAddressFormatValid(addr); !isValid {
		return &Result{Success, false, Unknown, ""}
	}

	addrType, err := e.Client.GetAccount(addr)
	if err != nil {
		return &Result{Failure, false, Unknown, err.Error()}
	}

	if addrType == Unknown {
		return &Result{Success, false, Unknown, ""}
	}

	return &Result{Success, true, addrType, ""}
}

// IsAddressFormatValid ...
func (e *IOST) IsAddressFormatValid(addr string) bool {
	re := regexp.MustCompile(`\A[a-z0-9_]{5,11}\z`)
	return re.MatchString(addr)
}
