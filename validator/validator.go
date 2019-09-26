package validator

type Validator interface {
	ValidateAddress(addr string, isTestnet bool) bool
}

func New(chainName string) Validator {
	switch chainName {
	case "bitcoin":
		return bitcoinvalidator.New()
	}
}
