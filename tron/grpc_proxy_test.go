package tron

import (
	"context"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"testing"
)

func TestNewGrpcProxy(t *testing.T) {
	proxy := NewGrpcProxy()
	defer proxy.Close()

	var err error

	var addr = "TCE9ifhQEMqBrqALokosKDh73mQsxEfxRZ"
	var ctx = context.Background()
	method := "/protocol.Wallet/GetAccountResource" // 获取账户能量
	in := new(core.Account)
	in.Address, err = common.DecodeCheck(addr)
	if err != nil {
		t.Fatalf("decode address err: %v", err)
	}

	var out = new(AccountResourceMessage)

	err = proxy.Invoke(ctx, method, in, out)
	if err != nil {
		t.Fatalf("invoke error: %v", err)
	}

	t.Logf("get account resource success: %v", out)
	t.Logf("TotalEnergyWeight: %d", out.TotalEnergyWeight)
	t.Logf("TotalNetLimit: %d", out.TotalNetLimit)
	t.Logf("EnergyLimit: %d", out.EnergyLimit) // 剩余能量
	t.Logf("TotalEnergyLimit: %d", out.TotalEnergyLimit)
	t.Logf("FreeNetUsed: %d", out.FreeNetUsed)
}
