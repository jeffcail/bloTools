package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

var (
	EthNetworkError                      = "the network ID for this client failed !!!"
	EthBlockTransactionSenderError       = "eth block transaction sender failed !!!"
	EthBlockTransactionReceiptError      = "the receipt of a transaction by transaction hash failed !!!"
	EthBlockSingleTransactionByHashError = "the transaction with the given hash failed !!!"
)

type Transaction struct {
	Hash     common.Hash     // 区块hash
	Value    *big.Int        // 传输的值，以 Wei 为单位
	Gas      uint64          // 发送者提供的燃料
	GasPrice *big.Int        // 发送者提供的燃料价格，以 Wei 为单位
	Nonce    uint64          // 交易顺序，防止重放攻击
	Data     []byte          // 数据
	To       *common.Address // 交易发送方
	From     common.Address  // 交易接收方
	Status   uint64          // 交易状态
}

// EthBlockTransaction 查询交易区块内交易记录
func (e *EthClient) EthBlockTransaction(blockNumber int64) []*Transaction {
	block := e.EthBlockContentByBlockNumber(blockNumber)

	transactions := make([]*Transaction, 0)

	for idx, tx := range block.Transactions() {
		transaction := &Transaction{
			Hash:     tx.Hash(),
			Value:    tx.Value(),
			Gas:      tx.Gas(),
			GasPrice: tx.GasPrice(),
			Nonce:    tx.Nonce(),
			Data:     tx.Data(),
			To:       tx.To(),
		}

		transaction.From = e.inquireFromByHash(context.Background(), tx, block.Hash(), uint(idx))
		transaction.Status = e.inquireTransactionReceipt(context.Background(), tx.Hash()).Status

		transactions = append(transactions, transaction)
	}
	return transactions
}

// EthBlockSingleTransactionByHash 单条交易
func (e *EthClient) EthBlockSingleTransactionByHash(hash string) (*Transaction, bool) {
	txHash := common.HexToHash(hash)
	tx, isPending, err := e.r.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatalf("【%s】%s", EthBlockSingleTransactionByHashError, err)
	}

	t := &Transaction{
		Hash:     tx.Hash(),
		Value:    tx.Value(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Nonce:    tx.Nonce(),
		Data:     tx.Data(),
		To:       tx.To(),
	}

	return t, isPending
}

// EthTransactionEth 交易eth
func (e *EthClient) EthTransactionEth(fromPrivateKey, toAddress string, value *big.Int, gasLimit uint64, data []byte) (bool, error) {
	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		return false, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return false, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := e.r.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return false, err
	}
	gasPrice, err := e.r.SuggestGasPrice(context.Background())
	if err != nil {
		return false, err
	}
	tAddress := common.HexToAddress(toAddress)
	tx := types.NewTransaction(nonce, tAddress, value, gasLimit, gasPrice, data)
	chainID, err := e.r.NetworkID(context.Background())
	if err != nil {
		return false, err
	}

	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return false, err
	}

	err = e.r.SendTransaction(context.Background(), signTx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// inquireFromByHash
func (e *EthClient) inquireFromByHash(ctx context.Context, tx *types.Transaction, hash common.Hash, idx uint) common.Address {
	sender, err := e.r.TransactionSender(ctx, tx, hash, idx)
	if err != nil {
		log.Fatalf("【%s】%s", EthBlockTransactionSenderError, err)
		return common.Address{}
	}
	return sender
}

// inquireFromByHash
func (e *EthClient) inquireTransactionReceipt(ctx context.Context, hash common.Hash) *types.Receipt {
	receipt, err := e.r.TransactionReceipt(ctx, hash)
	if err != nil {
		log.Fatalf("【%s】%s", EthBlockTransactionReceiptError, err)
		return nil
	}
	return receipt
}
