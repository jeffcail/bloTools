package tron

import (
	"encoding/json"
	"fmt"
	"github.com/jeffcail/gorequest"
	"net/http"
)

type (
	HttpClientInterface interface {
		GetLatestSignalTransaction(address, symbolAddress string) (*TransactionRes, error)
		GetTransactions(address, symbolAddress string) (*TransactionRes, error)
	}
)

type HttpClient struct {
	c   *http.Client
	url string
}

func NewHttpClient(url string) *HttpClient {
	return &HttpClient{
		c:   new(http.Client),
		url: url,
	}
}

func (h *HttpClient) GetLatestSignalTransaction(address, symbolAddress string) (*TransactionRes, error) {
	// https://api.trongrid.io/v1/accounts/%s/transactions/trc20?limit=1&contract_address=%s
	bytes, err := gorequest.Get(fmt.Sprintf(h.url, address, symbolAddress), nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(TransactionRes)
	if err = json.Unmarshal(bytes, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HttpClient) GetTransactions(address, symbolAddress string) (*TransactionRes, error) {
	// https://api.trongrid.io/v1/accounts/%s/transactions/trc20?contract_address=%s
	bytes, err := gorequest.Get(fmt.Sprintf(h.url, address, symbolAddress), nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(TransactionRes)
	if err = json.Unmarshal(bytes, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
