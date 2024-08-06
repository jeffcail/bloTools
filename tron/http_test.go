package tron

import (
	"fmt"
	"testing"
)

func TestNewHttpClient(t *testing.T) {

	tests := map[string]struct {
		test func(t *testing.T)
	}{}

	t.Parallel()

	for name, tt := range tests {
		t.Run(name, tt.test)
	}

}

func TestHttpClient_GetTrc10TokenPrecision(t *testing.T) {
	url := fmt.Sprintf("%s%s", "", "wallet/getassetissuelist")

	precision, err := NewHttpClient("").GetTrc10TokenPrecision(url)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	t.Logf("precision: %v", precision)
}
