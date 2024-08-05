package bsc

import "testing"

func TestNewBscAddress(t *testing.T) {
	var tests = map[string]struct {
		test func(t *testing.T)
	}{
		"GenerateBscAddress": {TestBscAddress_GenerateBscAddress},
	}

	t.Parallel()

	for name, tt := range tests {
		t.Run(name, tt.test)
	}
}

func TestBscAddress_GenerateBscAddress(t *testing.T) {
	privateKey, address := NewBscAddress().GenerateBscAddress()
	t.Logf("private key: %s", privateKey)
	t.Logf("address: %s", address)
}
