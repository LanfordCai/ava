package validator

import (
	"ava/httpclient"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// TeraClient ...
type TeraClient struct {
	Endpoint string
}

// GetAccount ...
func (c *TeraClient) GetAccount(addr string) (AddressType, error) {
	url := fmt.Sprintf("%s/api/v2/GetBalance", c.Endpoint)
	reqBody := map[string]string{
		"AccountID": addr,
	}
	resp, err := httpclient.Post(url, reqBody)
	if err != nil {
		return Unknown, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Unknown, err
	}

	if resp.StatusCode > 299 {
		return Unknown, nil
	}

	var r teraGetBalanceResp
	err = json.Unmarshal(data, &r)
	if err != nil {
		return Unknown, err
	}

	if r.Result == 1 {
		return Normal, nil
	}

	return Unknown, nil
}

type teraGetBalanceResp struct {
	Result int `json:"result"`
}
