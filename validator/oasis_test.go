package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOasisIsAddressFormatValid(t *testing.T) {
	validator := &Oasis{}

	var validCases = []string{
		"oasis1qqs6ylpfurhf6gc9mw232fkmrt3d0673lyzc5xf2",
		"oasis1qrad7s7nqm4gvyzr8yt2rdk0ref489rn3vn400d6",
		"oasis1qp29h8ykmxet46eqzw0wennrmmy4al3xzv37m3ca",
		"oasis1qp7hpak0wuymn3pqtwf2cpaxsm25v6ydhuzkxupe",
	}

	for _, addr := range validCases {
		assert.True(t, validator.IsAddressFormatValid(addr, Mainnet))
		assert.True(t, validator.IsAddressFormatValid(addr, Testnet))
	}

	var invalidCases = []string{
		"oasis2qqs6ylpfurhf6gc9mw232fkmrt3d0673lyzc5xf2",
		"Oasis1qqs6ylpfurhf6gc9mw232fkmrt3d0673lyzc5xf2",
		"OASIS1QQS6YLPFURHF6GC9MW232FKMRT3D0673LYZC5XF2",
		"oasis1qqs6ylpfurhf6gc9mw232fkmrt3d0673lyzc5xf3",
		"oasis1qqs6ylpfurhf6gc9mw232fkmrt3d0673lyzc5xf3qq",
		"oasis1qqs6ylprhf6gc9mw232fkmrt3d0673lyzc5xf3",
	}

	for _, addr := range invalidCases {
		assert.False(t, validator.IsAddressFormatValid(addr, Mainnet))
		assert.False(t, validator.IsAddressFormatValid(addr, Testnet))
	}
}
