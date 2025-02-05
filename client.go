package go1inch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	inchURL = "https://api.1inch.dev/swap/v5.2/"
)

type Network int64

const (
	Eth         Network = 1
	Bsc         Network = 56
	Matic       Network = 137
	Optimism    Network = 10
	Arbitrum    Network = 42161
	GnosisChain Network = 100
	Avalanche   Network = 43114
	Fantom      Network = 250
	Klaytn      Network = 8217
	Auror       Network = 1313161554
)

func NewDefaultClient() *Client {
	return NewClient("", http.DefaultClient)
}

func NewClient(apiKey string, httpClient *http.Client) *Client {
	return &Client{
		Http:   httpClient,
		apiKey: apiKey,
	}
}

func setQueryParam(endpoint *string, params []map[string]interface{}) {
	first := true
	for _, param := range params {
		for i := range param {
			if first {
				*endpoint = fmt.Sprintf("%s?%s=%v", *endpoint, i, param[i])
				first = false
			} else {
				*endpoint = fmt.Sprintf("%s&%s=%v", *endpoint, i, param[i])
			}
		}
	}
}

func (c *Client) doRequest(ctx context.Context, net Network, endpoint, method string, expRes interface{}, reqData interface{}, opts ...map[string]interface{}) (int, http.Header, error) {
	callURL := fmt.Sprintf("%s%s%s", inchURL, strconv.FormatInt(int64(net), 10), endpoint)

	var dataReq []byte
	var err error

	if reqData != nil {
		dataReq, err = json.Marshal(reqData)
		if err != nil {
			return 0, nil, err
		}
	}

	if len(opts) > 0 && len(opts[0]) > 0 {
		setQueryParam(&callURL, opts)
	}
	req, err := http.NewRequestWithContext(ctx, method, callURL, bytes.NewBuffer(dataReq))
	if err != nil {
		return 0, nil, err
	}

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Accept", "application/json")
	if len(c.apiKey) > 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}

	resp, err := c.Http.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		if expRes != nil {
			err = json.Unmarshal(body, expRes)
			if err != nil {
				return 0, resp.Header, err
			}
		}
		return resp.StatusCode, resp.Header, nil

	default:
		return resp.StatusCode, resp.Header, fmt.Errorf("%s", body)
	}
}
