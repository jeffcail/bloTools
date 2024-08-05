package eth

import (
	"fmt"
	"testing"
)

func TestEthClient_EthBlockNumber(t *testing.T) {

	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	blockNumber := cli.EthBlockNumber()
	fmt.Println(blockNumber)
}

func TestEthClient_EthBlockNumber2(t *testing.T) {
	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	blockNumber := cli.EthBlockNumber2()
	fmt.Println(blockNumber)
}

func TestEthClient_EthBlockContentByBlockNumber(t *testing.T) {
	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	blockNumber := 19707994
	block := cli.EthBlockContentByBlockNumber(int64(blockNumber))

	fmt.Println(block.Number())              // 区块高度
	fmt.Println(block.Time())                // 区块时间戳
	fmt.Println(block.Difficulty().Uint64()) // 区块难度
	fmt.Println(block.Hash().Hex())          // 区块hash
	fmt.Println(len(block.Transactions()))   // 区块交易计数
	fmt.Println(block.Transactions())        // 区块交易列表
}

func TestEthClient_EthBlockTransactionCountByBlockNumber(t *testing.T) {
	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	blockNumber := 19707994
	block := cli.EthBlockContentByBlockNumber(int64(blockNumber))

	fmt.Println(block.Hash())                                        // 0xd40639354dbe1614924f941f519fd7baee48717032910084a3167097d69e1010
	count := cli.EthBlockTransactionCountByBlockNumber(block.Hash()) // 0xd40639354dbe1614924f941f519fd7baee48717032910084a3167097d69e1010
	fmt.Println(count)                                               // 66
}
