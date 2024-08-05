package bsc

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/shopspring/decimal"
)

type (
	RpcClientInterface interface {
		Call(params interface{}) (string, error)
		LatestBlockNumber() (string, error)
		GetBlockByNumber(blockNumber string) (*BlockResponse, error)
		GetBlockByHash(blockHash string) (*BlockResponse, error)
		ChainID() (string, error)
		EstimateGas(callMsg ethereum.CallMsg) (string, error)
	}
)

type RpcClient struct {
	client *rpc.Client
}

func NewRpcClient(network string) *RpcClient {
	client, err := rpc.Dial(network)
	if err != nil {
		panic(err)
	}

	return &RpcClient{client: client}
}

// LatestBlockNumber 最新区块高度
func (rc *RpcClient) LatestBlockNumber() (string, error) {
	var response string
	err := rc.client.Call(&response, "eth_blockNumber")
	if err != nil {
		return "", nil
	}
	s := parseBscResponseString(response)
	newFromString, err := decimal.NewFromString(s)
	if err != nil {
		return "", err
	}
	return newFromString.String(), nil
}

// GetBlockByNumber get block by number
func (rc *RpcClient) GetBlockByNumber(blockNumber string) (*BlockResponse, error) {
	strHex := strToHex(blockNumber)
	response := new(BlockResponse)
	if err := rc.client.Call(response, "eth_getBlockByNumber", strHex, true); err != nil {
		return nil, err
	}
	return response, nil
}

// GetBlockByHash get block by hash
func (rc *RpcClient) GetBlockByHash(blockHash string) (*BlockResponse, error) {
	response := new(BlockResponse)
	if err := rc.client.Call(response, "eth_getBlockByHash", blockHash, true); err != nil {
		return nil, err
	}
	return response, nil
}

// ChainID chain id
func (rc *RpcClient) ChainID() (string, error) {
	var response string
	err := rc.client.Call(&response, "eth_chainId")
	if err != nil {
		return "", err
	}
	chanId := parseBscResponseString(response)
	return chanId, nil
}

// EstimateGas
// 执行并估算一个交易需要的gas用量。该次交易不会写入区块链。注意，由于多种原因，例如EVM的机制 及节点旳性能，估算的数值可能比实际用量大的多。
func (rc *RpcClient) EstimateGas(callMsg ethereum.CallMsg) (string, error) {
	var response string
	err := rc.client.Call(&response, "eth_estimateGas", toCallArg(callMsg))
	if err != nil {
		return "", err
	}
	data := parseBscResponseString(response)
	return data, nil
}

// Call 拨号
func (rc *RpcClient) Call(params interface{}) (string, error) {
	var response string
	err := rc.client.Call(&response, "eth_call", params, "latest")
	if err != nil {
		return "", err
	}
	return response, nil
}
