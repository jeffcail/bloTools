package tron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	bloTools "github.com/jeffcail/blcTools"
	"github.com/spf13/cast"
	"io"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
)

type (
	RpcClientInterface interface {
		Call(method string, params []interface{}) ([]byte, error)
		GetBlockHeight() (int64, error)
		GetBlockByNumber(number string) (*BlockByNumberResp, error)
		GetTransactionByHash(hash string) (*Transactions, error)
		GetTransactionReceipt(hash string) (*TransactionReceipt, error)
		GetTRXBalance(address string) (*big.Int, error)
		GetTrc20Balance(addr string, con string) (*big.Int, error)
		GetTokenName(token string) (string, error)
		GetSymbol(token string) (string, error)
		GetDecimal(token string) (string, error)
		EstimateGas(m map[string]string) (*big.Int, error)
	}
)

// GetBlockHeight get block height
func (r *RpcClient) GetBlockHeight() (int64, error) {
	body, err := r.Call("eth_blockNumber", []interface{}{})
	if err != nil {
		return 0, err
	}
	var resp struct {
		Result string `json:"result"`
		Error  string `json:"error,omitempty"`
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return 0, err
	}
	count, err := bloTools.HexToEthereumNumber(resp.Result)
	if err != nil {
		return 0, err
	}
	return count.Int64(), nil
}

// GetBlockByNumber get block info by number
func (r *RpcClient) GetBlockByNumber(number string) (*BlockByNumberResp, error) {
	body, err := r.Call("eth_getBlockByNumber", []interface{}{number, true})
	if err != nil {
		return nil, err
	}
	var resp struct {
		Jsonrpc string             `json:"jsonrpc,omitempty"`
		Id      string             `json:"id,omitempty"`
		Result  *BlockByNumberResp `json:"result,omitempty"`
		Error   string             `json:"error,omitempty"`
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return resp.Result, err
	}
	return resp.Result, nil
}

// GetTransactionByHash get transaction info by hash
func (r *RpcClient) GetTransactionByHash(hash string) (*Transactions, error) {
	body, err := r.Call("eth_getTransactionByHash", []interface{}{hash})
	if err != nil {
		return nil, err
	}
	var resp struct {
		Jsonrpc string        `json:"jsonrpc,omitempty"`
		Id      string        `json:"id,omitempty"`
		Result  *Transactions `json:"result,omitempty"`
		Error   string        `json:"error,omitempty"`
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return resp.Result, err
	}
	return resp.Result, nil
}

// GetTransactionReceipt get transaction receipt
func (r *RpcClient) GetTransactionReceipt(hash string) (*TransactionReceipt, error) {
	body, err := r.Call("eth_getTransactionReceipt", []interface{}{hash})
	if err != nil {
		return nil, err
	}
	var resp struct {
		Result *TransactionReceipt `json:"result,omitempty"`
		Error  string              `json:"error"`
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return nil, fmt.Errorf(resp.Error)
	}
	return resp.Result, nil
}

// GetTRXBalance get trx balance
func (r *RpcClient) GetTRXBalance(address string) (*big.Int, error) {
	if strings.HasPrefix(address, "T") {
		hex, err := tl.AddressB58ToHex(address)
		if err != nil {
			panic(err)
		}
		address = hex[2:]
	}
	body, err := r.Call("eth_getBalance", []interface{}{address, "latest"})
	if err != nil {
		return big.NewInt(0), err
	}
	resp := &Resp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return big.NewInt(0), err
	}
	if resp.Result == "" {
		return big.NewInt(0), fmt.Errorf(resp.Error.Message)
	}
	return bloTools.HexToBigInt(resp.Result), nil
}

// GetTrc20Balance get trc20 balance
func (r *RpcClient) GetTrc20Balance(addr string, con string) (*big.Int, error) {
	if strings.HasPrefix(con, "T") {
		hex, err := tl.AddressB58ToHex(con)
		if err != nil {
			panic(err)
		}
		con = hex[2:]
	}
	addrB, err := address.Base58ToAddress(addr)
	if err != nil {
		return nil, fmt.Errorf("invalid address %s: %v", addr, addr)
	}
	s := "0x70a082310000000000000000000000" + strings.ReplaceAll(addrB.Hex(), "0x", "")
	reqArgs := map[string]interface{}{
		"to":   con,
		"data": s,
	}
	call, err := r.Call("eth_call", []interface{}{reqArgs, "latest"})
	if err != nil {
		return nil, err
	}
	var res struct {
		Jsonrpc string `json:"jsonrpc"`
		Id      string `json:"id"`
		Result  string `json:"result"`
	}
	if err = json.Unmarshal(call, &res); err != nil {
		return nil, err
	}
	property := tl.parseErc20StringProperty(res.Result)
	return big.NewInt(cast.ToInt64(property)), nil
}

// GetTokenName get token name
func (r *RpcClient) GetTokenName(token string) (string, error) {
	if strings.HasPrefix(token, "T") {
		hex, err := tl.AddressB58ToHex(token)
		if err != nil {
			panic(err)
		}
		token = hex[2:]
	}
	request := map[string]interface{}{
		"to":   token,
		"data": "0x06fdde03",
	}
	call, err := r.Call("eth_call", []interface{}{request, "latest"})
	if err != nil {
		return "", err
	}
	var Res struct {
		Jsonrpc string `json:"jsonrpc"`
		Id      string `json:"id"`
		Result  string `json:"result"`
	}
	if err = json.Unmarshal(call, &Res); err != nil {
		return "", err
	}

	property := tl.parseErc20StringProperty(Res.Result)
	return property, nil
}

// GetSymbol get symbol
func (r *RpcClient) GetSymbol(token string) (string, error) {
	if strings.HasPrefix(token, "T") {
		hex, err := tl.AddressB58ToHex(token)
		if err != nil {
			panic(err)
		}
		token = hex[2:]
	}
	request := map[string]interface{}{
		"to":   token,
		"data": "0x95d89b41",
	}
	call, err := r.Call("eth_call", []interface{}{request, "latest"})
	if err != nil {
		return "", err
	}
	var res struct {
		Jsonrpc string `json:"jsonrpc"`
		Id      string `json:"id"`
		Result  string `json:"result"`
	}
	if err = json.Unmarshal(call, &res); err != nil {
		return "", err
	}
	property := tl.parseErc20StringProperty(res.Result)
	return property, nil
}

// GetDecimal get decimal
func (r *RpcClient) GetDecimal(token string) (string, error) {
	if strings.HasPrefix(token, "T") {
		hex, err := tl.AddressB58ToHex(token)
		if err != nil {
			panic(err)
		}
		token = hex[2:]
	}
	request := map[string]interface{}{
		"to":   token,
		"data": "0x313ce567",
	}
	call, err := r.Call("eth_call", []interface{}{request, "latest"})
	if err != nil {
		return "", err
	}
	var res struct {
		Jsonrpc string `json:"jsonrpc"`
		Id      string `json:"id"`
		Result  string `json:"result"`
	}
	if err = json.Unmarshal(call, &res); err != nil {
		return "", err
	}

	property := tl.parseErc20StringProperty(res.Result)
	return property, nil
}

// EstimateGas get pre estimate gas
func (r *RpcClient) EstimateGas(m map[string]string) (*big.Int, error) {
	body, err := r.Call("eth_estimateGas", []interface{}{m})
	if err != nil {
		return big.NewInt(0), err
	}
	resp := &Resp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return big.NewInt(0), err
	}
	if resp.Result == "" {
		return big.NewInt(0), fmt.Errorf(resp.Error.Message)
	}

	return bloTools.HexToBigInt(resp.Result), nil
}

type RpcClient struct {
	client  *http.Client
	network string
}

var co = int32(0)
var tl *TrTool

func init() {
	atomic.StoreInt32(&co, 0)
	tl = NewTronTool()
}

func (err *EthereumError) Error() string {
	return fmt.Sprintf("BLOCK-API RPC ERROR Ethereum %d %s", err.Code, err.Message)
}

type rpcCallRequest struct {
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      string        `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
}

func NewRPC(network string) *RpcClient {
	chain := &RpcClient{
		client:  new(http.Client),
		network: network,
	}
	return chain
}

// Call 拨号
func (r *RpcClient) Call(method string, params []interface{}) ([]byte, error) {
	data := rpcCallRequest{
		Method:  method,
		Params:  params,
		Id:      os.Getenv("TRON-NETWORK-ID"),
		JSONRPC: os.Getenv("TRON-JSON-RPC"),
	}
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", r.network, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
