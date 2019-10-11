package eos

import (
	"log"
	"os"
	"regexp"
)

// Validator - EOS address validator
type Validator struct {
	client            *Client
	contractWhitelist []string
}

func init() {
	if os.Getenv("AVA_EOS_BASE_URL") == "" {
		log.Panic("AVA_EOS_BASE_URL not found")
	}
}

// New - Create a new EOS address validator
func New(baseURL string, contractWhitelist []string) *Validator {
	client := Client{BaseURL: baseURL}
	return &Validator{client: &client, contractWhitelist: contractWhitelist}
}

// ValidateAddress ...
func (e *Validator) ValidateAddress(addr string, isTestnet bool) (isValid bool, msg string) {
	if !e.withValidFormat(addr) {
		return false, "Invalid format"
	}

	if !e.isRegistered(addr) {
		return false, "Unregistered"
	}

	if e.isContract(addr) {
		if contains(e.contractWhitelist, addr) {
			return true, "Whitelist contract"
		}
		return false, "contract isn't in whitelist"
	}
	return true, ""
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

func contains(sli []string, target string) bool {
	for _, s := range sli {
		if s == target {
			return true
		}
	}

	return false
}
