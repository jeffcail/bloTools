package tron

import "testing"

func TestNewGrpcClient(t *testing.T) {

	tests := map[string]struct {
		test func(t *testing.T)
	}{
		"GetAccountResource": {TestGrpcServerClient_GetAccountResource},
	}

	t.Parallel()

	for name, tt := range tests {
		t.Run(name, tt.test)
	}

}

func TestGrpcServerClient_GetAccountResource(t *testing.T) {
	client, err := NewGrpcClient("")
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	var addrss = ""
	resource, err := client.GetAccountResource(addrss)
	if err != nil {
		t.Fatalf("get account resource failed: %v", err)
	}
	t.Logf("get account resource success: %v", resource)
	t.Logf("TotalEnergyWeight: %d", resource.TotalEnergyWeight)
	t.Logf("TotalNetLimit: %d", resource.TotalNetLimit)
	t.Logf("EnergyLimit: %d", resource.EnergyLimit) // 剩余能量
	t.Logf("TotalEnergyLimit: %d", resource.TotalEnergyLimit)
	t.Logf("FreeNetUsed: %d", resource.FreeNetUsed)
}
