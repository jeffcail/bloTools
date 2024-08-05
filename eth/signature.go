package eth

import (
	"bytes"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

var (
	EthSignatureHexToECDSAError = "signature hex to ECDSA failed !!!"
	EthSignatureSignError       = "sign private key failed !!!"
	EthSignError                = "Sign failed !!!"
	EthVerifySignEcrecoverError = "the uncompressed public key that created the given signature failed !!!"
)

// EthSignature 签名
func (e *EthClient) EthSignature(privateKey string, data []byte) string {

	ecdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("【%s】%s", EthSignatureHexToECDSAError, err)
	}

	hash := crypto.Keccak256Hash(data)

	signature, err := crypto.Sign(hash.Bytes(), ecdsa)
	if err != nil {
		log.Fatalf("【%s】%s", EthSignatureSignError, err)
	}

	return hexutil.Encode(signature)

}

// VerifySignature 验证签名
func (e *EthClient) VerifySignature(privateKey string, data []byte) bool {
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("【%s】%s", EthSignatureHexToECDSAError, err)
	}

	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	hash := crypto.Keccak256Hash(data)
	sig, err := crypto.Sign(hash.Bytes(), ecdsaPrivateKey)
	if err != nil {
		log.Fatalf("【%s】%s", EthSignError, err)
	}

	signatureNoRecoverID := sig[:len(sig)-1] // remove recovery id
	return crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
}

// VerifySignatureOfEqual 验证签名
func (e *EthClient) VerifySignatureOfEqual(privateKey string, data []byte) bool {
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	hash := crypto.Keccak256Hash(data)
	sig, err := crypto.Sign(hash.Bytes(), ecdsaPrivateKey)
	if err != nil {
		log.Fatalf("【%s】%s", EthSignError, err)
	}

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		log.Fatalf("【%s】%s", EthVerifySignEcrecoverError, err)
	}

	return bytes.Equal(sigPublicKey, publicKeyBytes)
}

// VerifySignatureOfSigToPub 验证签名
func (e *EthClient) VerifySignatureOfSigToPub(privateKey string, data []byte) bool {
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	hash := crypto.Keccak256Hash(data)
	sig, err := crypto.Sign(hash.Bytes(), ecdsaPrivateKey)
	if err != nil {
		log.Fatalf("【%s】%s", EthSignError, err)
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	return bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
}
