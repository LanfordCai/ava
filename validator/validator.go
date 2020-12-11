package validator

type validator interface {
	ValidateAddress(addr string, network NetworkType) *Result
}
