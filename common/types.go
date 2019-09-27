package common

// ValidationResult - Address validation result
type ValidationResult struct {
	IsValid bool
	Msg     string
}

// NewValidationResult - Create a new ValidationResult
func NewValidationResult(isValid bool, msg string) ValidationResult {
	return ValidationResult{IsValid: isValid, Msg: msg}
}

// type Validator interface {
// 	ValidateAddress(addr string, isTestnet bool) bool
// }

// func New(chainName string) Validator {
// 	switch chainName {
// 	case "bitcoin":
// 		return bitcoinvalidator.New()
// 	}
// }
