package tron

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

type TrAddress struct{}

func NewTrAddress() *TrAddress {
	return &TrAddress{}
}

// GenerateTronAddress 生成私钥和地址
func (t *TrAddress) GenerateTronAddress() (privateKey, b58Address string) {
	ecdsaPrivateKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", ""
	}
	if len(ecdsaPrivateKey.D.Bytes()) != 32 {
		for {
			ecdsaPrivateKey, err = btcec.NewPrivateKey(btcec.S256())
			if err != nil {
				continue
			}
			if len(ecdsaPrivateKey.D.Bytes()) == 32 {
				break
			}
		}
	}
	b58Address = address.PubkeyToAddress(ecdsaPrivateKey.ToECDSA().PublicKey).String()
	privateKey = hex.EncodeToString(ecdsaPrivateKey.D.Bytes())
	return
}
