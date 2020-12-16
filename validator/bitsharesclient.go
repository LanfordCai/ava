package validator

import (
	"github.com/LanfordCai/ava/httpclient"
)

// BitsharesClient ...
type BitsharesClient struct {
	Endpoint string
}

// GetAccount ...
func (c *BitsharesClient) GetAccount(addr string) (AddressType, error) {
	reqBody := map[string]interface{}{
		"method":  "get_account",
		"params":  []string{addr},
		"jsonrpc": "2.0",
		"id":      1,
	}
	resp, err := httpclient.Post(c.Endpoint, reqBody)
	if err != nil {
		return Unknown, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return Unknown, nil
	}

	return Normal, nil
}
