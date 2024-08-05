package eth

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ClientDialError = "Dial eth client error"
)

type EthClient struct {
	r *ethclient.Client
}

func NewEthClient(network string) *EthClient {

	c, err := ethclient.Dial(network)
	if err != nil {
		panic(ClientDialError + "【😭】" + err.Error())
	}
	e := &EthClient{
		r: c,
	}

	return e
}
