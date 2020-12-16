package validator

import (
	"ava/httpclient"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// EOSClient ...
type EOSClient struct {
	Endpoint string
}

// GetAccount ...
func (c *EOSClient) GetAccount(addr string) (AddressType, error) {
	url := fmt.Sprintf("%s/v1/chain/get_account", c.Endpoint)
	reqBody := map[string]string{
		"account_name": addr,
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

// GetABIVersion ...
func (c *EOSClient) GetABIVersion(addr string) (string, error) {
	url := fmt.Sprintf("%s/v1/chain/get_abi", c.Endpoint)
	reqBody := map[string]string{
		"account_name": addr,
	}
	resp, err := httpclient.Post(url, reqBody)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return "", nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var r EOSGetABIResp
	err = json.Unmarshal(data, &r)
	if err != nil {
		return "", err
	}

	return r.Version, nil
}

// EOSGetABIResp ...
type EOSGetABIResp struct {
	Version string `json:"version"`
}
