//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what methodGen.queryVar 단위 테스트 (tx면 qtx, 아니면 server.Queries)

package ssac

import "testing"

func TestMethodGenQueryVar(t *testing.T) {
	if got := (&methodGen{UseTx: true}).queryVar(); got != "qtx" {
		t.Errorf("UseTx=true queryVar = %q, want %q", got, "qtx")
	}
	if got := (&methodGen{UseTx: false}).queryVar(); got != "server.Queries" {
		t.Errorf("UseTx=false queryVar = %q, want %q", got, "server.Queries")
	}
}
