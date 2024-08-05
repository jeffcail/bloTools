package eth

import (
	"fmt"
	"math/big"
	"testing"
)

func TestEthClient_EthBlockTransaction(t *testing.T) {
	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	var blockNumber int64 = 5671744
	transactions := cli.EthBlockTransaction(blockNumber)

	for ids, transaction := range transactions {
		fmt.Println(transaction.Hash) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(ids)              // 0
	}
}

func TestEthClient_EthBlockSingleTransactionByHash(t *testing.T) {
	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	var blockNumber int64 = 5671744
	transactions := cli.EthBlockTransaction(blockNumber)

	for _, transaction := range transactions {
		tx, isPending := cli.EthBlockSingleTransactionByHash(transaction.Hash.String())
		fmt.Println(isPending)
		fmt.Println(tx)
	}

}

func TestEthClient_EthTransactionEth(t *testing.T) {
	var mainNetWork = "https://rinkeby.infura.io"
	cli = NewEthClient(mainNetWork)
	fmt.Println("eth client dial success...😊😊😊")

	fromPrivateKey := "998539ea327486d88ee3c4fd03f172a28e9594812bfff82e8aad8409e35bdab5"
	//fromAddress := "0xE358F7AEB27114B02676C95Bb8C21b523551cADC"
	//toPrivateKey := "16ea985bb291b038d8f9d634a8b0a9f77db016ff47705a1b20ba9f3aea4b3932"
	toAddress := "0x3D5231c2608568056e0b085bB50195b25c5ed362"

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)
	var data []byte
	ok, err := cli.EthTransactionEth(fromPrivateKey, toAddress, value, gasLimit, data)
	fmt.Println(ok, err)
}
