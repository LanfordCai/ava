package ava

import (
	"strings"

	"github.com/LanfordCai/ava/bitcoin"
	"github.com/LanfordCai/ava/bytom"
	"github.com/LanfordCai/ava/eos"
	"github.com/LanfordCai/ava/ethereum"
	"github.com/LanfordCai/ava/litecoin"
	"github.com/LanfordCai/ava/neo"
	"github.com/LanfordCai/ava/qtum"
	"github.com/LanfordCai/ava/ripple"
	"github.com/LanfordCai/ava/sia"
	"github.com/LanfordCai/ava/zcash"
)

// Address ...
type Address struct {
	ChainName string
	Address   string
	IsTestnet bool
}

// IsValid ...
func (addr *Address) IsValid() bool {
	chainName := strings.ToLower(addr.ChainName)
	switch chainName {
	case "bitcoin", "bitcoincash", "bitcoinsv", "bitcoinabc", "superbitcoin":
		return bitcoin.IsValidAddress(addr.Address, addr.IsTestnet)
	case "litecoin":
		return litecoin.IsValidAddress(addr.Address, addr.IsTestnet)
	case "zcash":
		return zcash.IsValidAddress(addr.Address, addr.IsTestnet)
	case "qtum":
		return qtum.IsValidAddress(addr.Address, addr.IsTestnet)
	case "ethereum", "ethereumclassic", "vechain":
		return ethereum.IsValidAddress(addr.Address)
	case "neo":
		return neo.IsValidAddress(addr.Address)
	case "sia", "siacore", "siaclassic":
		return sia.IsValidAddress(addr.Address)
	case "bytom":
		return bytom.IsValidAddress(addr.Address, addr.IsTestnet)
	case "ripple":
		return ripple.IsValidAddress(addr.Address)
	case "eos":
		return eos.IsValidAddress(addr.Address)
	}
	return false
}

func (addr *Address) IsContractAddress() bool {
	return true
}
