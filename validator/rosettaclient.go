package validator

import (
	"ava/httpclient"
	"fmt"
)

// RosettaClient ...
type RosettaClient struct {
	Endpoint string
}

// GetAccount ...
func (c *RosettaClient) GetAccount(ni RosettaNetworkIdentifier, addr string) (AddressType, error) {
	url := fmt.Sprintf("%s/account/balance", c.Endpoint)
	reqBody := map[string]interface{}{
		"network_identifier": ni,
		"account_identifier": RosettaAccountIdentifier{Address: addr},
		"jsonrpc":            "2.0",
		"id":                 "1",
	}
	resp, err := httpclient.Post(url, reqBody)
	if err != nil {
		return Unknown, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return Unknown, nil
	}

	return Normal, nil
}
