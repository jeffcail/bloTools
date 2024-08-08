package tron

import "testing"

func TestTool(t *testing.T) {
	tests := map[string]struct {
		test func(t *testing.T)
	}{
		"AddressB58ToHex":     {TestTrTool_AddressB58ToHex},
		"AddressHexToB58":     {TestTrTool_AddressHexToB58},
		"ValidateTronAddress": {TestTrTool_ValidateTronAddress},
		"AccuracyPrivateKey":  {TestTrTool_AccuracyPrivateKey},
		"ParseInputData":      {TestTrTool_ParseInputData},
	}

	t.Parallel()
	for name, tt := range tests {
		t.Run(name, tt.test)
	}
}

var tl *TrTool

func init() {
	tl = NewTronTool()
}

func TestTrTool_AddressB58ToHex(t *testing.T) {
	var addr = ""
	hex, err := tl.AddressB58ToHex(addr)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	t.Logf("hex: %v", hex)
}

func TestTrTool_AddressHexToB58(t *testing.T) {
	var hex = ""
	b58 := tl.AddressHexToB58(hex)
	t.Logf("b58: %v", b58)
}

func TestTrTool_ValidateTronAddress(t *testing.T) {
	var add = ""
	err := tl.ValidateTronAddress(add)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	t.Log("address is valid!")
}

func TestTrTool_AccuracyPrivateKey(t *testing.T) {
	var address = ""
	var privateKey = ""
	if err := tl.AccuracyPrivateKey(privateKey, address); err != nil {
		t.Fatalf("err: %v", err)
	}

	t.Log("private key is valid!")
}

func TestTrTool_ParseInputData(t *testing.T) {
	var input = "0xa9059cbb00000000000000000000000082a89268afbc38ddd327d369866706bcc309808700000000000000000000000000000000000000000000000000000000000003e8"
	to, amount := tl.ParseInputData(input)
	b58 := tl.AddressHexToB58(to)
	t.Logf("to: %s", to)
	t.Logf("tob58: %s", b58)
	t.Logf("amount: %s", amount)
}
