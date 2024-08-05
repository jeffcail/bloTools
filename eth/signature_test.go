package eth

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"testing"
)

func TestEthClient_EthSignature(t *testing.T) {
	cli = NewEthClient(mainNetwork)

	fmt.Println("eth client dial success...😊😊😊")

	privateKey, _ := cli.EthAccount()
	t.Log(privateKey) // 3d736405317acf4683eff93b516e7ac75df4c09da6f71bb62a6e0d984279fd5f

	var data = []byte("test")
	signature := cli.EthSignature(privateKey, data)
	t.Log(signature) // 0xc75002a67fe7fdf946208692f407f4a8dcb596b2454d35e9569918095c164e1f58fb3aaa1c05bed90fdff5d794ba836b1e6f32ba603acdf25efdf95a723cb18500
}

func TestEthClient_VerifySignature2(t *testing.T) {
	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	privateKey := "3d736405317acf4683eff93b516e7ac75df4c09da6f71bb62a6e0d984279fd5f"
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	key := ecdsaPrivateKey.Public()
	fmt.Println(key)

	publicKeyECDSA, ok := key.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte("test")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex())

	sig, err := crypto.Sign(hash.Bytes(), ecdsaPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hexutil.Encode(sig))

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches)

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches)

	signatureNoRecoverID := sig[:len(sig)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) //
}

func TestEthClient_VerifySignature(t *testing.T) {
	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	privateKey := "3d736405317acf4683eff93b516e7ac75df4c09da6f71bb62a6e0d984279fd5f"
	data := []byte("test")
	ok := cli.VerifySignature(privateKey, data)
	t.Log(ok)
}

func TestEthClient_VerifySignatureOfEqual(t *testing.T) {
	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	privateKey := "3d736405317acf4683eff93b516e7ac75df4c09da6f71bb62a6e0d984279fd5f"
	data := []byte("test")
	ok := cli.VerifySignatureOfEqual(privateKey, data)
	t.Log(ok)
}

func TestEthClient_VerifySignatureOfSigToPub(t *testing.T) {
	cli = NewEthClient(mainNetwork)
	fmt.Println("eth client dial success...😊😊😊")

	privateKey := "3d736405317acf4683eff93b516e7ac75df4c09da6f71bb62a6e0d984279fd5f"
	data := []byte("test")
	ok := cli.VerifySignatureOfSigToPub(privateKey, data)
	t.Log(ok)
}
