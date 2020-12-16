package validator

import "regexp"

// Bitshares address validator
type Bitshares struct {
	Client *BitsharesClient
}

var _ OnchainValidator = (*Bitshares)(nil)

// ValidateAddress ...
func (e *Bitshares) ValidateAddress(addr string, network NetworkType) *Result {
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
func (e *Bitshares) IsAddressFormatValid(addr string) bool {
	re := regexp.MustCompile(`\A[a-z0-9][a-z0-9.-]{1,30}\z`)
	return re.MatchString(addr)
}
