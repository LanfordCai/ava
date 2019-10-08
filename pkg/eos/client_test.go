package eos

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccount(t *testing.T) {
	baseURL := os.Getenv("AVA_EOS_BASE_URL")
	client := Client{BaseURL: baseURL}

	t.Run("get unregistered account", func(t *testing.T) {
		resp, err := client.getAccount("helloworld6")
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("get registered account", func(t *testing.T) {
		resp, err := client.getAccount("huobideposit")
		assert.Nil(t, err)
		assert.Equal(t, "huobideposit", resp.AccountName)
	})
}

func TestGetABI(t *testing.T) {
	baseURL := os.Getenv("AVA_EOS_BASE_URL")
	client := Client{BaseURL: baseURL}

	t.Run("get unregistered account", func(t *testing.T) {
		resp, err := client.getABI("helloworld6")
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("get registered non-contract account", func(t *testing.T) {
		resp, err := client.getABI("pxneosincome")
		assert.Nil(t, err)
		assert.Equal(t, "pxneosincome", resp.AccountName)
		assert.Nil(t, resp.ABI)
	})

	t.Run("get registered contract account", func(t *testing.T) {
		resp, err := client.getABI("huobideposit")
		assert.Nil(t, err)
		assert.Equal(t, "huobideposit", resp.AccountName)
		assert.NotNil(t, resp.ABI)
	})
}
