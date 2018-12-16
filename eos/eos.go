package eos

import (
	"regexp"
)

// IsValidAddress ...
func IsValidAddress(address string) bool {
	matched, err := regexp.MatchString("[^.12345abcdefghijklmnopqrstuvwxyz]", address)
	if err != nil || len(address) > 12 || len(address) == 0 {
		return false
	}
	return !matched
}
