package aeternity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAddress(t *testing.T) {
	validator := New()
	var validAddresses = []string{
		"ak_2QkttUgEyPixKzqXkJ4LX7ugbRjwCDWPBT4p4M2r8brjxUxUYd",
		"ak_KHfXhF2J6VBt3sUgFygdbpEkWi6AKBkr9jNKUCHbpwwagzHUs",
		"ak_zvU8YQLagjcfng7Tg8yCdiZ1rpiWNp1PBn3vtUs44utSvbJVR",
	}

	var invalidAddresses = []string{
		"ak_4h3c6RH52R",
		"akzvU8YQLagjcfng7Tg8yCdiZ1rpiWNp1PBn3vtUs44utSvbJVR",
		"ak_zvU8YQLagjcfng7Tg8yCdiZ1rpiWNp1PBn3vtUs44utSvbJV",
		"aK_zvU8YQLagjcfng7Tg8yCdiZ1rpiWNp1PBn3vtUs44utSvbJVR",
		"th_2SYg4CBCZpCBTnYncSzE25sRUMN4uFDDKc11pMD7u73J6xaP1g",
		"",
	}

	for _, a := range validAddresses {
		result := validator.ValidateAddress(a, false)
		assert.Equal(t, true, result.IsValid)
	}

	for _, a := range invalidAddresses {
		result := validator.ValidateAddress(a, false)
		assert.Equal(t, false, result.IsValid)
	}
}
