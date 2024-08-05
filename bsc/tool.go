package bsc

import (
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
	"strings"
)

// parseBscResponseString
func parseBscResponseString(data string) string {

	if strings.HasPrefix(data, "0x") {
		data = data[2:]
	}
	if len(data) > 128 {
	} else if len(data) == 64 {
	} else {
		value, ok := new(big.Int).SetString(data, 16)
		if ok {
			return value.String()
		}
	}

	return ""
}

func blockNumberToHex(blockNumber string) string {
	s := strToHex(blockNumber)
	return fmt.Sprintf("%x%s", "0x", s)
}

// strToHex 字符转十六进制
func strToHex(s string) string {
	a, _ := strconv.Atoi(s)
	return strconv.FormatInt(int64(a), 16)
}

func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["input"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}
