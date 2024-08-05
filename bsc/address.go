package bsc

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type BscAddress struct{}

func NewBscAddress() *BscAddress {
	return &BscAddress{}
}

func (b *BscAddress) GenerateBscAddress() (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", ""
	}

	// generate address of use private key
	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	privateKeyStr := hexutil.Encode(crypto.FromECDSA(privateKey))

	return privateKeyStr, address
}
