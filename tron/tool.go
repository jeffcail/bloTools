package tron

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	bloTools "github.com/jeffcail/blcTools"
	"github.com/jeffcail/blcTools/eth"
	"math/big"
	"strings"
)

type TrTool struct{}

func NewTronTool() *TrTool {
	return &TrTool{}
}

func (t *TrTool) AddressB58ToHex(b58Address string) (string, error) {
	a, err := address.Base58ToAddress(b58Address)
	if err != nil {
		return "", err
	}
	d := strings.Replace(a.Hex(), "0x41", "0x", -1)
	return d, nil
}

func (t *TrTool) AddressHexToB58(hexAddress string) string {
	if hexAddress == "" {
		return ""
	}
	d := strings.Replace(hexAddress, "0x", "41", -1)
	return address.HexToAddress(d).String()
}

// ValidateTronAddress 验证波场地址的有效性
func (t *TrTool) ValidateTronAddress(address string) error {
	// 波场地址长度为 34 个字符，并以 "T" 开头
	if len(address) != 34 || address[0] != 'T' {
		return errors.New("invalid address length or prefix")
	}

	// Base58 解码地址
	decoded := base58.Decode(address)
	if len(decoded) != 25 {
		return errors.New("invalid address length after decoding")
	}

	// 校验版本号（波场地址的版本号为 0x41）
	if decoded[0] != 0x41 {
		return errors.New("invalid address version")
	}

	// 提取原始地址和校验和
	rawAddress := decoded[:len(decoded)-4]
	checksum := decoded[len(decoded)-4:]

	// 计算实际校验和
	actualChecksum := bloTools.Sha256Checksum(rawAddress)

	// 比较校验和
	if !bloTools.Equal(checksum, actualChecksum[:4]) {
		return errors.New("checksum mismatch")
	}

	return nil
}

// AccuracyPrivateKey 私钥准确性
func (t *TrTool) AccuracyPrivateKey(privateKey, address string) error {
	hexAddress, err := eth.NewEthTool().PrivateKeyToAddress(privateKey)
	if err != nil {
		return err
	}

	b58 := t.AddressHexToB58(hexAddress)

	return t.validateAddress(b58, address)
}

// validateAddress
func (t *TrTool) validateAddress(toAddress, address string) error {
	var ead string
	if address[:1] == "T" {
		ead = address
	} else {
		ead = t.AddressHexToB58(address)
	}

	if toAddress != ead {
		return fmt.Errorf("invalid address or private key")
	}
	return nil
}

func (t *TrTool) parseErc20StringProperty(data string) string {
	if strings.HasPrefix(data, "0x") {
		data = data[2:]
	}
	if len(data) > 128 {
		n := t.parseErc20NumericProperty(data[64:128])
		if n != nil {
			l := n.Uint64()
			if 2*int(l) <= len(data)-128 {
				b, err := hex.DecodeString(data[128 : 128+2*l])
				if err == nil {
					return string(b)
				}
			}
		}
	} else if len(data) == 64 {
		return t.parseErc20NumericProperty(data).String()
	} else {
		value, success := new(big.Int).SetString(string(data), 16)
		if success {
			return value.String()
		}
	}
	return ""
}
func (t *TrTool) parseErc20NumericProperty(data string) *big.Int {
	if strings.HasPrefix(data, "0x") {
		data = data[2:]
	}
	if len(data) == 64 {
		var n big.Int
		_, ok := n.SetString(data, 16)
		if ok {
			return &n
		}
	}
	return nil
}

func (t *TrTool) ParseTokenAbbrAndName(s string) string {
	ds, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(ds)
}

// ParseInputData
// 对于每笔交易，根据交易的输入数据（input 字段）解析出相应的 Token 信息。如果交易调用了智能合约，input 字段将包含方法签名和参数。
// TRC-20 代币的 transfer 方法通常以 0xa9059cbb 开头，其后是接收地址和转账金额。
func (t *TrTool) ParseInputData(input string) (to, amount string) {
	if len(input) == 0 && len(input) <= 10 {
		return
	}
	has := strings.HasPrefix(input, "0xa9059cbb")
	if has {
		// 解析 Token 转账的接收者地址和金额
		to = "0x" + input[34:74]
		amountBigInt := new(big.Int)
		amountBigInt.SetString(input[74:], 16)
		amount = amountBigInt.String()
	}
	return to, amount
}
