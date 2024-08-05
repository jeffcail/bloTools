package mnemonics_test

import (
	"fmt"
	"github.com/jeffcail/blcTools/mnemonics"
	"testing"
)

func TestGenerateMnemonic(t *testing.T) {
	tests := map[string]struct {
		test func(t *testing.T)
	}{
		"Generate24Mnemonic": {TestGenerate24Mnemonic},
		"Generate12Mnemonic": {TestGenerate12Mnemonic},
	}
	t.Parallel()
	for name, tt := range tests {
		t.Run(name, tt.test)
	}
}

func TestGenerate12Mnemonic(t *testing.T) {
	fmt.Printf("mnemonic: %s\n", mnemonics.Generate24Mnemonic())
}

func TestGenerate24Mnemonic(t *testing.T) {
	fmt.Printf("mnemonic: %s\n", mnemonics.Generate12Mnemonic())
}
