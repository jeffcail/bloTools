package eth

import (
	"context"
	"log"
	"math/big"
)

var (
	InquireEthBalanceErr = "Inquire eth balance failed !!!"
)

// EthBalanceByAddress 获取账户余额
func (e *EthClient) EthBalanceByAddress(address string) *big.Float {
	toAddress := hexToAddress(address)
	balanceAt, err := e.r.BalanceAt(context.Background(), toAddress, nil)
	if err != nil {
		log.Fatal(InquireEthBalanceErr + " || " + err.Error())
	}
	return balanceToWei(balanceAt)
}

// EthBalanceByBlockNumber 获取该区块时的账户余额
func (e *EthClient) EthBalanceByBlockNumber(address string, blockNumber int64) *big.Float {
	toAddress := hexToAddress(address)
	n := big.NewInt(blockNumber)
	at, err := e.r.BalanceAt(context.Background(), toAddress, n)
	if err != nil {
		log.Fatal(InquireEthBalanceErr + " || " + err.Error())
	}
	return balanceToWei(at)
}
