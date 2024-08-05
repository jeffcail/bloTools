package eth

import (
	"testing"
)

var (
	mainNetwork = "https://cloudflare-eth.com"
)

var cli *EthClient

func TestNewClientEth(t *testing.T) {
	cli = NewEthClient(mainNetwork)
}
