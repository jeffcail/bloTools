package eth

import (
	"testing"
)

func TestNewEthTool(t *testing.T) {
	var tests = map[string]struct {
		test func(t *testing.T)
	}{
		"PrivateKeyToAddress": {TestEtTool_PrivateKeyToAddress},
	}

	t.Parallel()

	for name, tt := range tests {
		t.Run(name, tt.test)
	}
}

var et *EtTool

func init() {
	et = NewEthTool()
}

func TestEtTool_PrivateKeyToAddress(t *testing.T) {
	var privateKey = ""
	hex, err := et.PrivateKeyToAddress(privateKey)
	if err != nil {
		t.Fatalf("err : %v", err)
	}

	t.Logf("address: %v", hex)
}
