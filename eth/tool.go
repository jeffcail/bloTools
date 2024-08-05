package eth

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math"
	"math/big"
	"regexp"
)

type EtTool struct{}

func NewEthTool() *EtTool {
	return &EtTool{}
}

// PrivateKeyToAddress private key to hex address
func (e *EtTool) PrivateKeyToAddress(privateHexKey string) (string, error) {
	var pk string
	if privateHexKey[:2] == "0x" {
		pk = privateHexKey[2:]
	} else {
		pk = privateHexKey
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return "", err
	}

	publicKey := ecdsaPrivateKey.Public()
	publicKeyEcdsa, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("invalid public key")
	}

	hex := crypto.PubkeyToAddress(*publicKeyEcdsa).Hex()
	return hex, err
}

// hexToAddress
func hexToAddress(s string) common.Address {
	return common.HexToAddress(s)
}

// balanceToWei
// 以太坊中的数字是使用尽可能小的单位来处理的，因为它们是定点精度，在ETH中它是wei。要读取ETH值，您必须做计算wei/10^18
func balanceToWei(balance *big.Int) *big.Float {
	b := new(big.Float)
	b.SetString(b.String())
	ethValue := new(big.Float).Quo(b, big.NewFloat(math.Pow10(18)))
	return ethValue
}

var (
	regComStr = "^0x[0-9a-fA-F]{40}$"
)

// ValidaEthAddress 校验地址是否有效
func ValidaEthAddress(address string) bool {
	return regexp.MustCompile(regComStr).MatchString(address)
}
