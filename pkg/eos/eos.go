package eos

import (
	"regexp"

	"github.com/LanfordCai/ava/internal/common"
	"github.com/LanfordCai/ava/internal/utils"
)

// Validator - EOS address validator
type Validator struct {
	client            *Client
	contractWhitelist []string
}

// New - Create a new EOS address validator
func New(baseURL string, contractWhitelist []string) *Validator {
	client := Client{BaseURL: baseURL}
	return &Validator{client: &client, contractWhitelist: contractWhitelist}
}

// ValidateAddress ...
func (e *Validator) ValidateAddress(addr string, isTestnet bool) common.ValidationResult {
	if e.withValidFormat(addr) && e.isRegistered(addr) {
		if e.isContract(addr) {
			if utils.Contains(e.contractWhitelist, addr) {
				return common.NewValidationResult(true, "whitelist contract")
			}
			return common.NewValidationResult(false, "contract which isn't in whitelist")
		}
		return common.NewValidationResult(true, "")
	}
	return common.NewValidationResult(false, "unregistered or invalid format")
}

func (e *Validator) isRegistered(addr string) bool {
	if _, err := e.client.getAccount(addr); err != nil {
		return false
	}

	return true
}

func (e *Validator) isContract(addr string) bool {
	abiResp, err := e.client.getABI(addr)
	if err != nil {
		return false
	}

	return abiResp.ABI != nil
}

func (e *Validator) withValidFormat(addr string) bool {
	matched, err := regexp.MatchString("[^.12345abcdefghijklmnopqrstuvwxyz]", addr)
	if err != nil || len(addr) > 12 || len(addr) == 0 {
		return false
	}
	return !matched
}
