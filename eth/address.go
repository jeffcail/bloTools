package eth

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	GenerateEthAddressError = "Generate eth address failed !!!"
	CanNotType              = "cannot assert type: publicKey is not of type *ecdsa.PublicKey"
)

// genEthAddress
func genEthAddress() (string, string) {
	var (
		privateKey *ecdsa.PrivateKey
		err        error
	)
	privateKey, err = crypto.GenerateKey()
	if err != nil {
		panic(GenerateEthAddressError + "【😭】" + err.Error())
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyStr := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(CanNotType)
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return privateKeyStr, address
}
