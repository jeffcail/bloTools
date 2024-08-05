package tron

import "testing"

func TestTrAddress(t *testing.T) {
	var tests = map[string]struct {
		test func(t *testing.T)
	}{
		"GenerateTronAddress": {TestTrAddress_GenerateTronAddress},
	}
	t.Parallel()

	for name, tt := range tests {
		t.Run(name, tt.test)
	}

}

var ta *TrAddress

func init() {
	ta = NewTrAddress()
}

func TestTrAddress_GenerateTronAddress(t *testing.T) {
	privateKey, b58Address := ta.GenerateTronAddress()
	t.Logf("private key: %s", privateKey)
	t.Logf("b58 address: %s", b58Address)
}
