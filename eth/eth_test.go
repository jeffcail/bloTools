package eth

import (
	"fmt"
	"testing"
)

func TestEthClient_EthAccount(t *testing.T) {
	cli = NewEthClient(mainNetwork)

	fmt.Println("eth client dial success...😊😊😊")
	fmt.Println(cli.r)

	privateKey, address := cli.EthAccount()
	fmt.Println(privateKey)
	fmt.Println(address)
}

func TestEthClient_ValidateEthAddressIsSmartContract(t *testing.T) {
	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	address := "0xb45D0616870aAf9AC4970f111592f85107E1C831"
	isSmartContract := cli.ValidateEthAddressIsSmartContract(address)
	fmt.Println(isSmartContract)

	address2 := "0xe41d2489571d322189246dafa5ebde1f4699f498"
	isSmartContract2 := cli.ValidateEthAddressIsSmartContract(address2)
	fmt.Println(isSmartContract2)
}
