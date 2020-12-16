package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// DefaultTimeout ...
const DefaultTimeout = 10 * time.Second

// Get ...
func Get(url string) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}

// Post ...
func Post(url string, reqbody interface{}) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	body, err := json.Marshal(reqbody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}
