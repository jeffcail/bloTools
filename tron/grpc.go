package tron

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/golang/protobuf/proto"
	bloTools "github.com/jeffcail/blcTools"
	"github.com/jinzhu/copier"
)

type GrpcServerInterface interface {
	GetAccountResource(address string) (*AccountResourceMessage, error)
	GetDelegateResourcesV2(fromAddress string) ([]*DelegatedResourceList, error)
	DelegateResourceEnergy(from, fromPrivateKey, to string, balance int64, lock bool, lockPeriod int64) (string, error)
	UnDelegateResource(from, fromPrivateKey, to string, balance int64, lock bool) (string, error)
	GetCanDelegatedEnergyMaxSize(fromAddress string) (int64, error)
	UnFreezeEnergy(from, fromPrivateKey, to string) (string, error)
}

type GrpcServerClient struct {
	network string
	c       *client.GrpcClient
}

func NewGrpcClient(network string) (*GrpcServerClient, error) {
	g := &GrpcServerClient{
		network: network,
	}

	c := client.NewGrpcClient(network)
	if err := c.Start(); err != nil {
		return nil, err
	}

	g.c = c

	return g, nil
}

// GetAccountResource get account resource 获取账户能量
func (g *GrpcServerClient) GetAccountResource(address string) (*AccountResourceMessage, error) {
	if g.c != nil {
		res, err := g.c.GetAccountResource(address)
		if err != nil {
			return nil, err
		}

		var a = new(AccountResourceMessage)
		if err = copier.Copy(a, res); err != nil {
			return nil, err
		}
		return a, nil
	}
	return nil, errors.New("no start")
}

// GetDelegateResourcesV2 get delegate resource v2
func (g *GrpcServerClient) GetDelegateResourcesV2(fromAddress string) ([]*DelegatedResourceList, error) {
	if g.c != nil {
		res, err := g.c.GetDelegatedResourcesV2(fromAddress)
		if err != nil {
			return nil, err
		}

		var list []*DelegatedResourceList
		for _, re := range res {
			for _, resource := range re.DelegatedResource {
				var d = new(DelegatedResourceList)
				if err = copier.Copy(d, resource); err != nil {
					return nil, err
				}
				list = append(list, d)
			}
		}
		return list, nil
	}
	return nil, errors.New("no start")
}

// DelegateResourceEnergy delegate resource energy 委派能量
func (g *GrpcServerClient) DelegateResourceEnergy(from, fromPrivateKey, to string, balance int64, lock bool, lockPeriod int64) (string, error) {
	if g.c != nil {
		tx, err := g.c.DelegateResource(from, to, core.ResourceCode_ENERGY, balance, lock, lockPeriod)
		if err != nil {
			return "", err
		}

		if tx.GetResult().GetCode() != api.Return_SUCCESS {
			return "", errors.New(bloTools.CompactStr("trx transaction failed,", string(tx.GetResult().GetMessage())))
		}

		fromPrivateKeyBytes, err := hex.DecodeString(fromPrivateKey)
		if err != nil {
			return "", err
		}

		sk, _ := btcec.PrivKeyFromBytes(fromPrivateKeyBytes)
		rawData, err := proto.Marshal(tx.Transaction.GetRawData())
		if err != nil {
			return "", err
		}
		hash := sha256.New()
		hash.Write(rawData)
		sum := hash.Sum(nil)
		sig, err := crypto.Sign(sum, sk.ToECDSA())
		if err != nil {
			return "", err
		}
		tx.Transaction.Signature = append(tx.Transaction.Signature, sig)

		result, err := g.c.Broadcast(tx.Transaction)
		if err != nil {
			return "", errors.New(bloTools.CompactStr("Broadcast transaction failed1,", err.Error()))
		}

		if result.Code != api.Return_SUCCESS {
			return "", errors.New(bloTools.CompactStr("Broadcast transaction failed2", string(tx.GetResult().GetMessage())))
		}
		return hex.EncodeToString(tx.Txid), nil
	}
	return "", errors.New("no start")
}

// UnDelegateResource un delegate resource 取消能量委派
func (g *GrpcServerClient) UnDelegateResource(from, fromPrivateKey, to string, balance int64, lock bool) (string, error) {
	if g.c != nil {
		tx, err := g.c.UnDelegateResource(from, to, core.ResourceCode_ENERGY, balance, lock)
		if err != nil {
			return "", err
		}

		if tx.GetResult().GetCode() != api.Return_SUCCESS {
			return "", errors.New("trx transaction failed," + string(tx.GetResult().GetMessage()))
		}

		fromPrivateKeyBytes, err := hex.DecodeString(fromPrivateKey)
		if err != nil {
			return "", err
		}
		sk, _ := btcec.PrivKeyFromBytes(fromPrivateKeyBytes)
		rowData, err := proto.Marshal(tx.Transaction.GetRawData())
		if err != nil {
			return "", err
		}
		h256h := sha256.New()
		h256h.Write(rowData)
		hash := h256h.Sum(nil)
		signature, err := crypto.Sign(hash, sk.ToECDSA())
		if err != nil {
			return "", err
		}
		tx.Transaction.Signature = append(tx.Transaction.Signature, signature)
		result, err := g.c.Broadcast(tx.Transaction)
		if err != nil {
			return "", err
		}
		if result.Code != api.Return_SUCCESS {
			return "", errors.New("Broadcast transaction failed," + string(tx.GetResult().GetMessage()))
		}
	}
	return "", errors.New("no start")
}

// GetCanDelegatedEnergyMaxSize get can delegated energy max size
func (g *GrpcServerClient) GetCanDelegatedEnergyMaxSize(fromAddress string) (int64, error) {
	if g.c != nil {
		res, err := g.c.GetCanDelegatedMaxSize(fromAddress, 1)
		if err != nil {
			return 0, err
		}
		return res.GetMaxSize(), nil
	}
	return 0, errors.New("not start")
}

// UnFreezeEnergy un freeze energy
func (g *GrpcServerClient) UnFreezeEnergy(from, fromPrivateKey, to string) (string, error) {
	if g.c != nil {
		tx, err := g.c.UnfreezeBalance(from, to, core.ResourceCode_ENERGY)
		if err != nil {
			return "", err
		}
		if tx.GetResult().GetCode() != api.Return_SUCCESS {
			return "", errors.New(bloTools.CompactStr("trx transaction failed,", string(tx.GetResult().GetMessage())))
		}

		fromPrivateKeyBytes, err := hex.DecodeString(fromPrivateKey)
		if err != nil {
			return "", err
		}

		sk, _ := btcec.PrivKeyFromBytes(fromPrivateKeyBytes)
		rawData, err := proto.Marshal(tx.Transaction.GetRawData())
		if err != nil {
			return "", err
		}
		hash := sha256.New()
		hash.Write(rawData)
		sum := hash.Sum(nil)
		sig, err := crypto.Sign(sum, sk.ToECDSA())
		if err != nil {
			return "", err
		}
		tx.Transaction.Signature = append(tx.Transaction.Signature, sig)

		result, err := g.c.Broadcast(tx.Transaction)
		if err != nil {
			return "", err
		}
		if result.Code != api.Return_SUCCESS {
			return "", errors.New(bloTools.CompactStr("Broadcast transaction failed1,", string(result.GetMessage())))
		}
		return hex.EncodeToString(tx.Txid), nil
	}
	return "", errors.New("no start")
}
