package eth

import (
	"fmt"
	"testing"
)

func TestEthClient_EthBalance(t *testing.T) {
	cli = NewEthClient(mainNetwork)

	fmt.Println("eth client dial success...😊😊😊")
	//fmt.Println(cli.r)

	address := "0xb45D0616870aAf9AC4970f111592f85107E1C831"
	balance := cli.EthBalanceByAddress(address)
	fmt.Println(balance)
}

func TestEthClient_EthBalanceByBlockNumber(t *testing.T) {
	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	address := "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"
	blockNumber := 5532993
	balance := cli.EthBalanceByBlockNumber(address, int64(blockNumber))
	fmt.Println(balance)
}
