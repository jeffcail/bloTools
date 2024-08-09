package tron

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	bloTools "github.com/jeffcail/blcTools"
	"github.com/jeffcail/gorequest"
	"net/http"
	"strings"
)

type (
	HttpClientInterface interface {
		GetLatestSignalTransaction(address, symbolAddress string) (*TransactionRes, error)
		GetTransactions(address, symbolAddress string) (*TransactionRes, error)
		GetTrc10TokenPrecision(url string) (*GetTrc10TokenPrecision, error)
		IdentifyTransactionToken(rpcClient *RpcClient, transactionHash string) (string, string, string, string, error)
		GetTrc10Token(assetID string) (string, error)
	}
)

func (h *HttpClient) GetTrc10TokenPrecision(url string) (*GetTrc10TokenPrecision, error) {
	header := make(map[string]string)
	header["accept"] = "application/json"
	res, err := gorequest.Get(url, header, nil)
	if err != nil {
		return nil, err
	}
	var data = new(GetTrc10TokenPrecision)
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

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

// IdentifyTransactionToken
// 识别交易通证
// response: token, contractAddress, transactionType, assetID, err
func (h *HttpClient) IdentifyTransactionToken(rpcClient *RpcClient, transactionHash string) (string, string, string, string, error) {
	hash := transactionHash[2:]
	url := bloTools.CompactStr(h.url, "wallet/gettransactionbyid")

	var header = make(map[string]string)
	var p = make(map[string]interface{})
	header["accept"] = "application/json"
	header["content-type"] = "application/json"
	p["value"] = hash

	res, err := gorequest.Post(url, header, p)
	if err != nil {
		return "", "", "", "", errors.New(fmt.Sprintf("根据交易Hash【%v】获取交易信息失败 err:【%v】", hash, err))
	}
	t := new(GetTransactionByID)
	_ = json.Unmarshal(res, &t)
	if t.RawData.Contract[0].Type == "TransferContract" { // 无合约 默认TRX
		return "TRX", "", "", "", nil
	}
	if t.RawData.Contract[0].Type == "TriggerSmartContract" { // 有合约 获取此交易记录的币种 TRC20 USDT
		if strings.HasPrefix(t.RawData.Contract[0].Parameter.Value.Data, "a9059cbb") {
			contractAddress := t.RawData.Contract[0].Parameter.Value.ContractAddress
			token, err := rpcClient.GetSymbol(contractAddress)
			if err != nil {
				return "", "", "", "", err
			}
			return token, contractAddress, "TriggerSmartContract", "", nil
		}
	}
	if t.RawData.Contract[0].Type == "TransferAssetContract" { // 无合约 TRC10
		assetID, _ := hex.DecodeString(t.RawData.Contract[0].Parameter.Value.AssetName)
		token, err := h.GetTrc10Token(string(assetID))
		if err != nil {
			return "", "", "", string(assetID), err
		}
		return token, "", "TransferAssetContract", "", nil
	}
	return "other", "", "", "", nil
}

// GetTrc10Token
// TRC 10 获取币种
func (h *HttpClient) GetTrc10Token(assetID string) (string, error) {
	url := bloTools.CompactStr(h.url, "wallet/getassetissuebyid")
	header := make(map[string]string)
	header["accept"] = "application/json"
	header["content-type"] = "application/json"
	p := make(map[string]interface{})
	p["value"] = assetID

	res, err := gorequest.Post(url, header, p)
	if err != nil {
		return "", errors.New(fmt.Sprintf("TRC10 获取交易记录币种失败 err: %v", err))
	}
	asset := &GetAssetIssueByID{}
	_ = json.Unmarshal(res, &asset)
	token, err := hex.DecodeString(asset.Abbr)
	if err != nil {
		return "", errors.New(fmt.Sprintf("TRC10 解析交易记录币种失败 err: %v", err))
	}
	return string(token), nil
}
