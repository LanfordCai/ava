package validator

import (
	"regexp"
)

// EOS address validator
type EOS struct {
	Client *EOSClient
}

var _ OnchainValidator = (*EOS)(nil)
var _ ContractChecker = (*EOS)(nil)

// ValidateAddress ...
func (e *EOS) ValidateAddress(addr string, network NetworkType) *Result {
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
func (e *EOS) IsAddressFormatValid(addr string) bool {
	re := regexp.MustCompile(`\A[.a-z1-5]{1,12}\z`)
	return re.MatchString(addr)
}

// IsContractDeployed ...
func (e *EOS) IsContractDeployed(addr string) (bool, error) {
	version, err := e.Client.GetABIVersion(addr)
	if err != nil {
		return false, err
	}

	if version != "" {
		return true, nil
	}

	return false, nil
}
