package validator

import "strings"

// Validator ...
type Validator interface {
	ValidateAddress(addr string, network NetworkType) *Result
}

// Prefixer ...
type Prefixer interface {
	GetPrefix(network NetworkType) string
}

// IsPrefixValid ...
func IsPrefixValid(v Prefixer, addr string, network NetworkType) bool {
	return strings.HasPrefix(addr, v.GetPrefix(network))
}

// AddressWithoutPrefix ...
func AddressWithoutPrefix(v Prefixer, addr string, network NetworkType) string {
	return strings.TrimPrefix(addr, v.GetPrefix(network))
}
