package eos

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Client - EOS client
type Client struct {
	BaseURL string
}

// GetAccountResp - response of get account
type GetAccountResp struct {
	AccountName string `json:"account_name"`
}

// GetABIResp - response of get abi
type GetABIResp struct {
	AccountName string      `json:"account_name"`
	ABI         interface{} `json:"abi"`
}

func (c *Client) getABI(account string) (*GetABIResp, error) {
	body := map[string]string{"account_name": account}
	var resp GetABIResp
	err := c.call("chain", "get_abi", body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) getAccount(account string) (*GetAccountResp, error) {
	body := map[string]string{"account_name": account}
	var resp GetAccountResp
	err := c.call("chain", "get_account", body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) call(baseAPI string, endpoint string, body interface{}, out interface{}) error {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/v1/%s/%s", c.BaseURL, baseAPI, endpoint)
	log.Print(url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Print(err)
		return err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	log.Print(string(respBody))
	if err != nil {
		return err
	}

	if resp.StatusCode > 299 {
		return errors.New(string(respBody))
	}

	err = json.Unmarshal(respBody, out)

	if err != nil {
		return err
	}

	return nil
}
