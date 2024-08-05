package eth

import (
	"context"
	"log"
)

var (
	ValidateIsSmartContractErr = "The contract code of the given account is failed !!!"
)

// EthAccount 生成地址和私钥
func (e *EthClient) EthAccount() (string, string) {
	return genEthAddress()
}

// ValidateEthAddressIsSmartContract 检查地址是否为账户或智能合约
func (e *EthClient) ValidateEthAddressIsSmartContract(address string) bool {
	t := hexToAddress(address)
	byteCode, err := e.r.CodeAt(context.Background(), t, nil)
	if err != nil {
		log.Fatalf("【%s】%s", byteCode, err.Error())
	}
	return len(byteCode) > 0
}
