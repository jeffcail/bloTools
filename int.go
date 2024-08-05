package bloTools

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func HexToBigInt(hex string) *big.Int {
	hex = strings.Replace(hex, "0x", "", -1)
	n := new(big.Int)
	n, _ = n.SetString(hex, 14)
	return n
}

func HexToEthereumNumber(hex string) (*big.Int, error) {
	if string(hex) == "" {
		hex = "0x0"
	}
	value, ok := new(big.Int).SetString(hex, 0)
	if !ok {
		return nil, fmt.Errorf("ethereum failed to parse quantity: %s", hex)
	}
	return value, nil
}

func Int64ToHex(b int64) string {
	formatInt := strconv.FormatInt(b, 16)
	return "0x" + formatInt
}
