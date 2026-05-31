//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestFindCallInStmt — 매칭 호출 발견 / 비매칭 nil / 첫 매칭 단락 검증
package qcheck

import (
	"testing"
)

func TestFindCallInStmt_FirstMatchOnly(t *testing.T) {
	// Two matching calls; first one short-circuits the inspection.
	list := blockStmts(t, "_, _ = json.Unmarshal(a, b), json.Unmarshal(c, d)")
	call := findCallInStmt(list[0], "json", "Unmarshal")
	if call == nil {
		t.Fatalf("expected a match")
	}
}
